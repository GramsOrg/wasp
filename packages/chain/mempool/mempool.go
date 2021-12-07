// mempool implements a buffer of requests sent to the ISCP chain, essentially a backlog of requests
// It contains both on-ledger and off-ledger requests. The mempool consists of 2 parts: the in-buffer and the pool
// All incoming requests are stored into the in-buffer first. Then they are asynchronously validated
// and moved to the pool itself.
package mempool

import (
	"bytes"
	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/chain"
	"sync"
	"time"

	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/iscp/rotate"
	"github.com/iotaledger/wasp/packages/metrics"
	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/packages/vm/core/blocklog"
)

type mempool struct {
	inBuffer                map[iscp.RequestID]iscp.RequestData
	inMutex                 sync.RWMutex
	poolMutex               sync.RWMutex
	inBufCounter            int
	outBufCounter           int
	inPoolCounter           int
	outPoolCounter          int
	stateReader             state.OptimisticStateReader
	pool                    map[iscp.RequestID]*requestRef
	chStop                  chan struct{}
	solidificationLoopDelay time.Duration
	log                     *logger.Logger
	mempoolMetrics          metrics.MempoolMetrics
}

var _ Mempool = &mempool{}

type requestRef struct {
	req          iscp.RequestData
	whenReceived time.Time
}

const (
	defaultSolidificationLoopDelay = 200 * time.Millisecond
	moveToPoolLoopDelay            = 20 * time.Millisecond
)

func New(stateReader state.OptimisticStateReader, log *logger.Logger, mempoolMetrics metrics.MempoolMetrics, solidificationLoopDelay ...time.Duration) Mempool {
	ret := &mempool{
		inBuffer:       make(map[iscp.RequestID]iscp.RequestData),
		stateReader:    stateReader,
		pool:           make(map[iscp.RequestID]*requestRef),
		chStop:         make(chan struct{}),
		log:            log.Named("m"),
		mempoolMetrics: mempoolMetrics,
	}
	if len(solidificationLoopDelay) > 0 {
		ret.solidificationLoopDelay = solidificationLoopDelay[0]
	} else {
		ret.solidificationLoopDelay = defaultSolidificationLoopDelay
	}
	go ret.moveToPoolLoop()
	return ret
}

func (m *mempool) addToInBuffer(req iscp.RequestData) bool {
	// just check if it is already in the pool
	if m.HasRequest(req.ID()) {
		return false
	}
	m.inMutex.Lock()
	defer m.inMutex.Unlock()
	// may be repeating but does not matter
	m.inBuffer[req.ID()] = req
	m.inBufCounter++
	return true
}

func (m *mempool) removeFromInBuffer(req iscp.Request) {
	m.inMutex.Lock()
	defer m.inMutex.Unlock()
	if _, ok := m.inBuffer[req.ID()]; ok {
		delete(m.inBuffer, req.ID())
		m.outBufCounter++
	}
}

// fills up the buffer with requests from the in-buffer
func (m *mempool) takeInBuffer(buf []iscp.RequestData) []iscp.RequestData {
	buf = buf[:0]
	m.inMutex.RLock()
	defer m.inMutex.RUnlock()

	for _, req := range m.inBuffer {
		buf = append(buf, req)
	}
	return buf
}

// addToPool adds request to the pool. It may fail
// returns true if it must be removed from the input buffer
func (m *mempool) addToPool(req iscp.RequestData) bool {
	reqid := req.ID()

	// checking in the state if request is processed. Reading may fail
	m.stateReader.SetBaseline()
	alreadyProcessed, err := blocklog.IsRequestProcessed(m.stateReader.KVStoreReader(), &reqid)
	if err != nil {
		// may be invalidated state. Do not remove from in-buffer yet
		m.log.Debugf("addToPool, IsRequestProcessed error: %v", err)
		return false
	}
	if alreadyProcessed {
		// remove from the in-buffer but not include into the pool
		return true
	}
	m.poolMutex.Lock()
	defer m.poolMutex.Unlock()

	if _, inPool := m.pool[reqid]; inPool {
		// already there, remove from the in-buffer
		return true
	}

	// put the request to the pool
	currentTime := time.Now()
	m.inPoolCounter++

	m.traceIn(req)

	m.pool[reqid] = &requestRef{
		req:          req,
		whenReceived: currentTime,
	}

	// return true to remove from the in-buffer
	return true
}

func (m *mempool) countRequestInMetrics(req iscp.RequestData) {
	// TODO refactor, this should be part of metrics logic.
	if req.IsOffLedger() {
		m.mempoolMetrics.CountOffLedgerRequestIn()
	} else {
		m.mempoolMetrics.CountOnLedgerRequestIn()
	}
}

// ReceiveRequests places requests into the inBuffer. InBuffer is unordered and non-deterministic
func (m *mempool) ReceiveRequests(reqs ...iscp.RequestData) {
	for _, req := range reqs {
		m.countRequestInMetrics(req)
		m.addToInBuffer(req)
	}
}

// ReceiveRequest receives a single request and returns whether that request has been added to the in-buffer
func (m *mempool) ReceiveRequest(req iscp.RequestData) bool {
	m.countRequestInMetrics(req)
	return m.addToInBuffer(req)
}

// RemoveRequests removes requests from the pool
func (m *mempool) RemoveRequests(reqs ...iscp.RequestID) {
	m.poolMutex.Lock()
	defer m.poolMutex.Unlock()

	for _, rid := range reqs {
		if _, ok := m.pool[rid]; !ok {
			continue
		}
		m.outPoolCounter++
		m.mempoolMetrics.CountRequestOut()
		m.mempoolMetrics.CountBlocksPerChain()
		elapsed := time.Since(m.pool[rid].whenReceived)
		m.mempoolMetrics.RecordRequestProcessingTime(rid, elapsed)
		delete(m.pool, rid)
		m.traceOut(rid)
	}
}

const traceInOut = false

func (m *mempool) traceIn(req iscp.RequestData) {
	rotateStr := ""
	if rotate.IsRotateStateControllerRequest(req) {
		rotateStr = "(rotate) "
	}
	logFn := m.log.Debugf
	if traceInOut {
		logFn = m.log.Infof
	}
	var timeLockTime time.Time
	var timeLockMilestone uint32

	if !req.IsOffLedger() {
		td := req.Unwrap().UTXO().Features().TimeLock()
		if td != nil {
			timeLockTime = td.Time
			timeLockMilestone = td.MilestoneIndex
		}
	}

	if !timeLockTime.IsZero() || timeLockMilestone > 0 {
		logFn("IN MEMPOOL %s%s (+%d / -%d) timelocked for %v until milestone %d",
			rotateStr, req.ID(), m.inPoolCounter, m.outPoolCounter, time.Until(timeLockTime), timeLockMilestone)
	} else {
		logFn("IN MEMPOOL %s%s (+%d / -%d)", rotateStr, req.ID(), m.inPoolCounter, m.outPoolCounter)
	}
}

func (m *mempool) traceOut(reqid iscp.RequestID) {
	if traceInOut {
		m.log.Infof("OUT MEMPOOL %s (+%d / -%d)", reqid, m.inPoolCounter, m.outPoolCounter)
	} else {
		m.log.Debugf("OUT MEMPOOL %s (+%d / -%d)", reqid, m.inPoolCounter, m.outPoolCounter)
	}
}

// don't process any request which deadline will expire within 10 minutes
const FallbackDeadlineMinAllowedInterval = time.Minute * 10

func isUnlockable(ref *requestRef, currentTime time.Time) bool {
	r := ref.req.(*iscp.OnLedgerRequestData)
	expiry, _ := r.Expiry()

	windowFrom := currentTime.Add(-FallbackDeadlineMinAllowedInterval)
	windowTo := currentTime.Add(FallbackDeadlineMinAllowedInterval)

	if expiry.Time.After(windowFrom) && expiry.Time.Before(windowTo) {
		return false
	}

	output, _ := ref.req.Unwrap().UTXO().Output().(iotago.TransIndepIdentOutput)

	unlockable := output.UnlockableBy(ref.req.SenderAddress(), &iotago.ExternalUnlockParameters{
		ConfUnix: uint64(currentTime.Unix()),
	})

	return unlockable
}

// isRequestReady for requests with paramsReady, the result is strictly deterministic
func isRequestReady(ref *requestRef, currentTime time.Time) (isReady, shouldBeRemoved bool) {
	if ref.req.IsOffLedger() {
		return true, false
	}

	r := ref.req.(*iscp.OnLedgerRequestData)

	// Skip anything with return amounts in this version.
	if _, ok := r.UTXO().Features().ReturnAmount(); ok {
		return false, true
	}

	if !isUnlockable(ref, currentTime) {
		return false, true
	}

	// time lock
	return r.TimeLock().Time.IsZero() || r.TimeLock().Time.Before(currentTime), false
}

// ReadyNow returns preliminary batch of requests for consensus.
// Note that later status of request may change due to the time change and time constraints
// If there's at least one committee rotation request in the mempool, the ReadyNow returns
// batch with only one request, the oldest committee rotation request
func (m *mempool) ReadyNow(currentTime ...time.Time) []iscp.RequestData {
	m.poolMutex.RLock()

	timeToValidate := time.Now()
	if len(currentTime) > 0 {
		timeToValidate = currentTime[0]
	}
	var oldestRotate iscp.RequestData
	var oldestRotateTime time.Time

	toRemove := []iscp.RequestID{}

	ret := make([]iscp.RequestData, 0, len(m.pool))
	for _, ref := range m.pool {
		rdy, shouldBeRemoved := isRequestReady(ref, timeToValidate)
		if shouldBeRemoved {
			toRemove = append(toRemove, ref.req.ID())
			continue
		}
		if !rdy {
			continue
		}
		ret = append(ret, ref.req)
		if !rotate.IsRotateStateControllerRequest(ref.req) {
			continue
		}
		// selecting oldest rotate request
		if oldestRotate == nil {
			oldestRotate = ref.req
			oldestRotateTime = ref.whenReceived
		} else {
			switch {
			case ref.whenReceived.Before(oldestRotateTime):
				oldestRotate = ref.req
				oldestRotateTime = ref.whenReceived
			case ref.whenReceived.Equal(oldestRotateTime):
				// for full determinism we take inti account not only time but also the request id
				if bytes.Compare(ref.req.ID().Bytes(), oldestRotate.ID().Bytes()) < 0 {
					oldestRotate = ref.req
					oldestRotateTime = ref.whenReceived
				}
			}
		}
	}
	m.poolMutex.RUnlock()
	go m.RemoveRequests(toRemove...)

	if oldestRotate != nil {
		return []iscp.RequestData{oldestRotate}
	}
	return ret
}

// ReadyFromIDs if successful, function returns a deterministic list of requests for running on the VM
// - (a list of missing requests), false if some requests not arrived to the mempool yet. For retry later
// - (a list of processable requests), true if the list can be deterministically calculated
// Note that (a list of processable requests) can be empty if none satisfies currentTime time constraint (timelock, fallback)
// For requests which are known and solidified, the result is deterministic
func (m *mempool) ReadyFromIDs(currentTime time.Time, reqIDs ...iscp.RequestID) ([]iscp.RequestData, []int, bool) {
	requests := make([]iscp.RequestData, 0, len(reqIDs))
	missingRequestIndexes := []int{}
	toRemove := []iscp.RequestID{}
	m.poolMutex.RLock()
	for i, reqID := range reqIDs {
		reqref, ok := m.pool[reqID]
		if !ok {
			missingRequestIndexes = append(missingRequestIndexes, i)
			continue
		}
		rdy, shouldBeRemoved := isRequestReady(reqref, currentTime)
		if rdy {
			requests = append(requests, reqref.req)
			continue
		}
		if shouldBeRemoved {
			toRemove = append(toRemove, reqref.req.ID())
		}
	}
	m.poolMutex.RUnlock()

	go m.RemoveRequests(toRemove...)

	return requests, missingRequestIndexes, len(missingRequestIndexes) == 0
}

// HasRequest checks if the request is in the pool
func (m *mempool) HasRequest(id iscp.RequestID) bool {
	m.poolMutex.RLock()
	defer m.poolMutex.RUnlock()

	_, ok := m.pool[id]
	return ok
}

func (m *mempool) GetRequest(id iscp.RequestID) iscp.Request {
	m.poolMutex.RLock()
	defer m.poolMutex.RUnlock()

	if reqRef, ok := m.pool[id]; ok {
		return reqRef.req
	}
	return nil
}

const waitRequestInPoolTimeoutDefault = 2 * time.Second

// WaitRequestInPool waits until the request appears in the pool but no longer than timeout
func (m *mempool) WaitRequestInPool(reqid iscp.RequestID, timeout ...time.Duration) bool {
	currentTime := time.Now()
	deadline := currentTime.Add(waitRequestInPoolTimeoutDefault)
	if len(timeout) > 0 {
		deadline = currentTime.Add(timeout[0])
	}
	for {
		if m.HasRequest(reqid) {
			return true
		}
		time.Sleep(10 * time.Millisecond)
		if time.Now().After(deadline) {
			return false
		}
	}
}

func (m *mempool) inBufferLen() int {
	m.inMutex.RLock()
	defer m.inMutex.RUnlock()
	return len(m.inBuffer)
}

const waitInBufferEmptyTimeoutDefault = 5 * time.Second

// WaitAllRequestsIn waits until in buffer becomes empty. Used in synchronous situations when the caller
// want to be sure all requests were fed into the pool. May create nondeterminism when used from goroutines
func (m *mempool) WaitInBufferEmpty(timeout ...time.Duration) bool {
	currentTime := time.Now()
	deadline := currentTime.Add(waitInBufferEmptyTimeoutDefault)
	if len(timeout) > 0 {
		deadline = currentTime.Add(timeout[0])
	}
	for {
		if m.inBufferLen() == 0 {
			return true
		}
		time.Sleep(10 * time.Millisecond)
		if time.Now().After(deadline) {
			return false
		}
	}
}

// Stats collects mempool stats
func (m *mempool) Info() chain.MempoolInfo {
	m.poolMutex.RLock()
	defer m.poolMutex.RUnlock()

	ret := chain.MempoolInfo{
		InPoolCounter:  m.inPoolCounter,
		OutPoolCounter: m.outPoolCounter,
		InBufCounter:   m.inBufCounter,
		OutBufCounter:  m.outBufCounter,
		TotalPool:      len(m.pool),
	}
	currentTime := time.Now()
	for _, ref := range m.pool {
		rdy, _ := isRequestReady(ref, currentTime)
		if rdy {
			ret.ReadyCounter++
		}
	}
	return ret
}

func (m *mempool) Close() {
	close(m.chStop)
}

// the loop validates and moves request from inBuffer to the pool
func (m *mempool) moveToPoolLoop() {
	buf := make([]iscp.RequestData, 0, 100)
	for {
		select {
		case <-m.chStop:
			return
		case <-time.After(moveToPoolLoopDelay):
			buf = m.takeInBuffer(buf)
			if len(buf) == 0 {
				continue
			}
			for i, req := range buf {
				if m.addToPool(req) {
					m.removeFromInBuffer(req)
				}
				buf[i] = nil // to please GC
			}
		}
	}
}
