package request

import (
	"net/http"
	"testing"
	"time"

	"github.com/iotaledger/wasp/packages/chain"
	"github.com/iotaledger/wasp/packages/isc"
	util "github.com/iotaledger/wasp/packages/testutil"
	"github.com/iotaledger/wasp/packages/util/expiringcache"
	"github.com/iotaledger/wasp/packages/webapi/v1/httperrors"
	"github.com/iotaledger/wasp/packages/webapi/v1/model"
	"github.com/iotaledger/wasp/packages/webapi/v1/routes"
	"github.com/iotaledger/wasp/packages/webapi/v1/testutil"
)

func shouldBeProcessedMocked(ret error) shouldBeProcessedFn {
	return func(_ chain.ChainCore, _ isc.OffLedgerRequest) error {
		return ret
	}
}

func newMockedAPI() *offLedgerReqAPI {
	return &offLedgerReqAPI{
		getChain: func(chainID isc.ChainID) chain.Chain {
			return &testutil.MockChain{}
		},
		shouldBeProcessed: shouldBeProcessedMocked(nil),
		requestsCache:     expiringcache.New(10 * time.Second),
	}
}

func testRequest(t *testing.T, instance *offLedgerReqAPI, chainID isc.ChainID, body interface{}, expectedStatus int) {
	testutil.CallWebAPIRequestHandler(
		t,
		instance.handleNewRequest,
		http.MethodPost,
		routes.NewRequest(":chainID"),
		map[string]string{"chainID": chainID.String()},
		body,
		nil,
		expectedStatus,
	)
}

func TestNewRequestBase64(t *testing.T) {
	instance := newMockedAPI()
	chainID := isc.RandomChainID()
	body := model.OffLedgerRequestBody{Request: model.NewBytes(util.DummyOffledgerRequest(chainID).Bytes())}
	testRequest(t, instance, chainID, body, http.StatusAccepted)
}

func TestNewRequestBinary(t *testing.T) {
	instance := newMockedAPI()
	chainID := isc.RandomChainID()
	body := util.DummyOffledgerRequest(chainID).Bytes()
	testRequest(t, instance, chainID, body, http.StatusAccepted)
}

func TestRequestAlreadyProcessed(t *testing.T) {
	instance := newMockedAPI()
	instance.shouldBeProcessed = shouldBeProcessedMocked(httperrors.BadRequest(""))

	chainID := isc.RandomChainID()
	body := util.DummyOffledgerRequest(chainID).Bytes()
	testRequest(t, instance, chainID, body, http.StatusBadRequest)
}

func TestWrongChainID(t *testing.T) {
	instance := newMockedAPI()
	body := util.DummyOffledgerRequest(isc.RandomChainID()).Bytes()
	testRequest(t, instance, isc.RandomChainID(), body, http.StatusBadRequest)
}
