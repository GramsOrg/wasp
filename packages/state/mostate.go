package state

import (
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/iscp/coreutil"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/buffered"
	"time"
)

// region mustOptimisticVirtualStateAccess ////////////////////////////////

// MustOptimisticVirtualState is a virtual state wrapper with global state baseline
// Once baseline is invalidated globally any subsequent access to the mustOptimisticVirtualStateAccess
// will lead to panic(coreutil.ErrorStateInvalidated)
type mustOptimisticVirtualStateAccess struct {
	state    VirtualStateAccess
	baseline coreutil.StateBaseline
}

// WrapMustOptimisticVirtualStateAccess wraps virtual state with state baseline in on object
// Does not copy buffers
func WrapMustOptimisticVirtualStateAccess(state VirtualStateAccess, baseline coreutil.StateBaseline) VirtualStateAccess {
	return &mustOptimisticVirtualStateAccess{
		state:    state,
		baseline: baseline,
	}
}

func (s *mustOptimisticVirtualStateAccess) BlockIndex() uint32 {
	s.baseline.MustValidate()
	defer s.baseline.MustValidate()

	return s.state.BlockIndex()
}

func (s *mustOptimisticVirtualStateAccess) Timestamp() time.Time {
	s.baseline.MustValidate()
	defer s.baseline.MustValidate()

	return s.state.Timestamp()
}

func (s *mustOptimisticVirtualStateAccess) PreviousStateHash() hashing.HashValue {
	s.baseline.MustValidate()
	defer s.baseline.MustValidate()

	return s.state.PreviousStateHash()
}

func (s *mustOptimisticVirtualStateAccess) StateCommitment() Commitment {
	s.baseline.MustValidate()
	defer s.baseline.MustValidate()

	return s.state.StateCommitment()
}

func (s *mustOptimisticVirtualStateAccess) KVStoreReader() kv.KVStoreReader {
	s.baseline.MustValidate()
	defer s.baseline.MustValidate()

	return s.state.KVStoreReader()
}

func (s *mustOptimisticVirtualStateAccess) ApplyStateUpdate(upd StateUpdate) {
	s.baseline.MustValidate()
	defer s.baseline.MustValidate()

	s.state.ApplyStateUpdate(upd...)
}

func (s *mustOptimisticVirtualStateAccess) ApplyBlock(block Block) error {
	s.baseline.MustValidate()
	defer s.baseline.MustValidate()

	return s.state.ApplyBlock(block)
}

func (s *mustOptimisticVirtualStateAccess) ExtractBlock() (Block, error) {
	s.baseline.MustValidate()
	defer s.baseline.MustValidate()

	return s.state.ExtractBlock()
}

func (s *mustOptimisticVirtualStateAccess) Commit(blocks ...Block) error {
	s.baseline.MustValidate()
	defer s.baseline.MustValidate()

	return s.state.Commit(blocks...)
}

func (s *mustOptimisticVirtualStateAccess) KVStore() *buffered.BufferedKVStoreAccess {
	s.baseline.MustValidate()
	defer s.baseline.MustValidate()

	return s.state.KVStore()
}

func (s *mustOptimisticVirtualStateAccess) Copy() VirtualStateAccess {
	s.baseline.MustValidate()
	defer s.baseline.MustValidate()

	return s.state.Copy()
}

func (s *mustOptimisticVirtualStateAccess) DangerouslyConvertToString() string {
	s.baseline.MustValidate()
	defer s.baseline.MustValidate()

	return s.state.DangerouslyConvertToString()
}
