package tests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/iotaledger/wasp/packages/cryptolib"
	"github.com/iotaledger/wasp/packages/parameters"
	"github.com/iotaledger/wasp/packages/vm/core/governance"
	"github.com/iotaledger/wasp/packages/vm/vmtypes"
	"github.com/iotaledger/wasp/tools/cluster/templates"
)

func TestWaspCLIExternalRotationGovAccessNodes(t *testing.T) {
	addAccessNode := func(w *WaspCLITest, pubKey string) {
		out := w.MustRun("chain", "gov-change-access-nodes", "accept", pubKey)
		out = w.GetReceiptFromRunPostRequestOutput(out)
		require.Regexp(t, `.*Error: \(empty\).*`, strings.Join(out, ""))
	}
	testWaspCLIExternalRotation(t, addAccessNode)
}

func TestWaspCLIExternalRotationPermitionlessAccessNodes(t *testing.T) {
	addAccessNode := func(w *WaspCLITest, pubKey string) {
		w.MustRun("chain", "access-nodes", "add", pubKey)
	}
	testWaspCLIExternalRotation(t, addAccessNode)
}

func testWaspCLIExternalRotation(t *testing.T, addAccessNode func(*WaspCLITest, string)) {
	// this test starts a chain on cluster of 4 nodes,
	// adds 1 new node as an access node (this node will be part of the new committee, this way it is synced)
	// then puts the chain on maintenance mode, stops the cluster
	// starts a new 4 nodes cluster (including the previous access node), runs the DKG on the new nodes,
	// rotates the chain state controller to the new cluster
	// stops the maintenance and ensure the chain is up-and-running

	w := newWaspCLITest(t, waspClusterOpts{
		nNodes:  4,
		dirName: "wasp-cluster-initial",
	})

	inccounterSCName := "inccounter"
	checkCounter := func(wTest *WaspCLITest, n int) {
		// test chain call-view command
		out := wTest.MustRun("chain", "call-view", inccounterSCName, "getCounter")
		out = wTest.MustPipe(out, "decode", "string", "counter", "int")
		require.Regexp(t, fmt.Sprintf(`(?m)counter:\s+%d$`, n), out[0])
	}

	committee, quorum := w.CommitteeConfig()
	out := w.MustRun(
		"chain",
		"deploy",
		"--chain=chain1",
		committee,
		quorum,
		fmt.Sprintf("--gov-controller=%s", w.WaspCliAddress.Bech32(parameters.L1().Protocol.Bech32HRP)),
	)
	chainID := regexp.MustCompile(`(.*)ChainID:\s*([a-zA-Z0-9_]*),`).FindStringSubmatch(out[len(out)-1])[2]

	// start a new wasp cluster
	w2 := newWaspCLITest(t, waspClusterOpts{
		nNodes:  4,
		dirName: "wasp-cluster-new-gov",
		modifyConfig: func(nodeIndex int, configParams templates.WaspConfigParams) templates.WaspConfigParams {
			// avoid port conflicts when running everything on localhost
			configParams.APIPort += 100
			configParams.DashboardPort += 100
			configParams.MetricsPort += 100
			configParams.NanomsgPort += 100
			configParams.PeeringPort += 100
			configParams.ProfilingPort += 100
			return configParams
		},
	})

	// adds node #0 from cluster2 as access node of the chain
	{
		node0peerInfo, err := w2.Cluster.WaspClient(0).GetPeeringSelf()
		require.NoError(t, err)

		// set trust relations between node0 of cluster 2 and all nodes of cluster 1
		w.Cluster.AddTrustedNode(node0peerInfo)
		cluster1PubKeys := make([]*cryptolib.PublicKey, len(w.Cluster.AllNodes()))
		for _, nodeIndex := range w.Cluster.Config.AllNodes() {
			// equivalent of "wasp-cli peer info"
			peerInfo, err := w.Cluster.WaspClient(nodeIndex).GetPeeringSelf()
			require.NoError(t, err)
			w2.MustRun("peering", "trust", peerInfo.PubKey, peerInfo.NetID)
			cluster1PubKeys[nodeIndex], err = cryptolib.NewPublicKeyFromString(peerInfo.PubKey)
			require.NoError(t, err)
		}

		// add node 0 from cluster 2 as an access node in the governance contract
		pubKey, err := cryptolib.NewPublicKeyFromString(node0peerInfo.PubKey)
		require.NoError(t, err)

		addAccessNode(w, pubKey.String())
	}

	// activate the chain on the new nodes
	w2.MustRun("chain", "add", "chain1", chainID)
	w2.MustRun("set", "chain", "chain1")
	w2.MustRun("chain", "activate")

	// deploy a contract, test its working
	{
		vmtype := vmtypes.WasmTime
		w.CopyFile(srcFile)

		// test chain deploy-contract command
		w.MustRun("chain", "deploy-contract", vmtype, inccounterSCName, "inccounter SC", file,
			"string", "counter", "int64", "42",
		)

		checkCounter(w, 42)
	}

	// init maintenance
	out = w.PostRequestGetReceipt("governance", "startMaintenance")
	require.Regexp(t, `.*Error: \(empty\).*`, strings.Join(out, ""))

	// check that node0 from clust2 is synced and maintenance is on
	for i := 0; ; i++ {
		if i >= 30 {
			t.Fatalf("Timeout waiting access node to be synched, last out=%v", out)
		}
		time.Sleep(1 * time.Second)
		var err error
		out, err = w2.Run("chain", "call-view", governance.Contract.Name, governance.ViewGetMaintenanceStatus.Name)
		if err != nil {
			t.Logf("Warning: call failed to ViewGetMaintenanceStatus: %v", err)
			continue
		}
		out, err = w2.Pipe(out, "decode", "string", governance.VarMaintenanceStatus, "bool")
		if err != nil {
			t.Logf("Warning: call failed to ViewGetMaintenanceStatus: %v", err)
			continue
		}
		if strings.Contains(out[0], "true") {
			break
		}
	}

	// stop the initial cluster
	w.Cluster.Stop()

	// run DKG on the new cluster, obtain the new state controller address
	out = w2.MustRun("chain", "rundkg")
	newStateControllerAddr := regexp.MustCompile(`(.*):\s*([a-zA-Z0-9_]*)$`).FindStringSubmatch(out[0])[2]

	// issue a governance rotatation via CLI
	out = w.MustRun("chain", "rotate", newStateControllerAddr)
	require.Regexp(t, `.*successfully.*`, strings.Join(out, ""))

	// stop maintenance
	// set the new nodes as the default (so querying the receipt doesn't fail)
	w.MustRun("set", "wasp.0.api", w2.Cluster.Config.APIHost(0))
	out = w.PostRequestGetReceipt("governance", "stopMaintenance")
	require.Regexp(t, `.*Error: \(empty\).*`, strings.Join(out, ""))

	// chain still works
	w2.MustRun("chain", "post-request", "-s", inccounterSCName, "increment")
	checkCounter(w2, 43)
}
