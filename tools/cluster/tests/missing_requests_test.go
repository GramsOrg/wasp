package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/iotaledger/wasp/client/chainclient"
	"github.com/iotaledger/wasp/contracts/native/inccounter"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/tools/cluster/templates"
)

func TestMissingRequests(t *testing.T) {
	// disable offledger request gossip between nodes
	modifyConfig := func(nodeIndex int, configParams templates.WaspConfigParams) templates.WaspConfigParams {
		configParams.OffledgerBroadcastUpToNPeers = 0
		return configParams
	}
	clu := newCluster(t, waspClusterOpts{nNodes: 4, modifyConfig: modifyConfig})
	cmt := []int{0, 1, 2, 3}
	threshold := uint16(4)
	addr, err := clu.RunDKG(cmt, threshold)
	require.NoError(t, err)

	chain, err := clu.DeployChain("chain", clu.Config.AllNodes(), cmt, threshold, addr)
	require.NoError(t, err)
	chainID := chain.ChainID

	chEnv := newChainEnv(t, clu, chain)
	chEnv.deployNativeIncCounterSC()

	waitUntil(t, chEnv.contractIsDeployed(), clu.Config.AllNodes(), 30*time.Second)

	userWallet, _, err := chEnv.Clu.NewKeyPairWithFunds()
	require.NoError(t, err)

	// deposit funds before sending the off-ledger request
	chClient := chainclient.New(clu.L1Client(), clu.WaspClient(0), chainID, userWallet)
	reqTx, err := chClient.DepositFunds(100)
	require.NoError(t, err)
	_, err = chain.CommitteeMultiClient().WaitUntilAllRequestsProcessedSuccessfully(chainID, reqTx, 30*time.Second)
	require.NoError(t, err)

	// send off-ledger request to all nodes except #3
	req := isc.NewOffLedgerRequest(chainID, nativeIncCounterSCHname, inccounter.FuncIncCounter.Hname(), dict.Dict{}, 0).Sign(userWallet)

	err = clu.WaspClient(0).PostOffLedgerRequest(chainID, req)
	require.NoError(t, err)
	err = clu.WaspClient(1).PostOffLedgerRequest(chainID, req)
	require.NoError(t, err)

	//------
	// send a dummy request to node #3, so that it proposes a batch and the consensus hang is broken
	req2 := isc.NewOffLedgerRequest(chainID, isc.Hn("foo"), isc.Hn("bar"), nil, 1).Sign(userWallet)
	err = clu.WaspClient(3).PostOffLedgerRequest(chainID, req2)
	require.NoError(t, err)
	//-------

	// expect request to be successful, as node #3 must ask for the missing request from other nodes
	waitUntil(t, chEnv.counterEquals(43), clu.Config.AllNodes(), 30*time.Second)
}
