// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package admapi

import (
	"github.com/pangpanglabs/echoswagger/v2"

	"github.com/iotaledger/wasp/packages/authentication"
	"github.com/iotaledger/wasp/packages/authentication/shared/permissions"
	"github.com/iotaledger/wasp/packages/chains"
	"github.com/iotaledger/wasp/packages/dkg"
	"github.com/iotaledger/wasp/packages/metrics/nodeconnmetrics"
	"github.com/iotaledger/wasp/packages/peering"
	"github.com/iotaledger/wasp/packages/registry"
	"github.com/iotaledger/wasp/packages/users"
)

func AddEndpoints(
	adm echoswagger.ApiGroup,
	network peering.NetworkProvider,
	tnm peering.TrustedNetworkManager,
	userManager *users.UserManager,
	chainRecordRegistryProvider registry.ChainRecordRegistryProvider,
	dkShareRegistryProvider registry.DKShareRegistryProvider,
	nodeIdentityProvider registry.NodeIdentityProvider,
	chainsProvider chains.Provider,
	nodeProvider dkg.NodeProvider,
	shutdownFunc ShutdownFunc,
	nodeConnectionMetrics nodeconnmetrics.NodeConnectionMetrics,
	authConfig authentication.AuthConfiguration,
	nodeOwnerAddresses []string,
) {
	claimValidator := func(claims *authentication.WaspClaims) bool {
		// The API will be accessible if the token has an 'API' claim
		return claims.HasPermission(permissions.API)
	}

	authentication.AddV1Authentication(adm.EchoGroup(), userManager, nodeIdentityProvider, authConfig, claimValidator)
	addShutdownEndpoint(adm, shutdownFunc)
	addNodeOwnerEndpoints(adm, nodeIdentityProvider, nodeOwnerAddresses)
	addChainRecordEndpoints(adm, chainRecordRegistryProvider, chainsProvider)
	addChainMetricsEndpoints(adm, chainsProvider)
	addChainEndpoints(adm, &chainWebAPI{
		chainRecordRegistryProvider: chainRecordRegistryProvider,
		dkShareRegistryProvider:     dkShareRegistryProvider,
		nodeIdentityProvider:        nodeIdentityProvider,
		chains:                      chainsProvider,
		network:                     network,
		// TODO: what happened to the metrics?
		nodeConnectionMetrics: nodeConnectionMetrics,
	})
	addDKSharesEndpoints(adm, dkShareRegistryProvider, nodeProvider)
	addPeeringEndpoints(adm, chainRecordRegistryProvider, network, tnm)
	addAccessNodesEndpoints(adm, chainRecordRegistryProvider, tnm)
}
