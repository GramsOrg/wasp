package services

import (
	"github.com/iotaledger/wasp/packages/chains"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/webapi/v2/dto"
	"github.com/iotaledger/wasp/packages/webapi/v2/interfaces"
	"github.com/iotaledger/wasp/packages/webapi/v2/models"
)

type MetricsService struct {
	chainProvider chains.Provider
}

func NewMetricsService(chainProvider chains.Provider) interfaces.MetricsService {
	return &MetricsService{
		chainProvider: chainProvider,
	}
}

func (c *MetricsService) GetAllChainsMetrics() *dto.ChainMetrics {
	chain := c.chainProvider()
	if chain == nil {
		return nil
	}

	metrics := chain.GetNodeConnectionMetrics()
	registered := metrics.GetRegistered()

	return &dto.ChainMetrics{
		InAliasOutput:      dto.MapMetricItem(metrics.GetInAliasOutput()),
		InOnLedgerRequest:  dto.MapMetricItem(metrics.GetInOnLedgerRequest()),
		InOutput:           dto.MapMetricItem(metrics.GetInOutput()),
		InStateOutput:      dto.MapMetricItem(metrics.GetInStateOutput()),
		InTxInclusionState: dto.MapMetricItem(metrics.GetInTxInclusionState()),
		InMilestone:        dto.MapMetricItem(metrics.GetInMilestone()),

		OutPublishGovernanceTransaction: dto.MapMetricItem(metrics.GetOutPublishGovernanceTransaction()),
		OutPullLatestOutput:             dto.MapMetricItem(metrics.GetOutPullLatestOutput()),
		OutPullOutputByID:               dto.MapMetricItem(metrics.GetOutPullOutputByID()),
		OutPullTxInclusionState:         dto.MapMetricItem(metrics.GetOutPullTxInclusionState()),
		OutPublisherStateTransaction:    dto.MapMetricItem(metrics.GetOutPublishStateTransaction()),

		RegisteredChainIDs: registered,
	}
}

func (c *MetricsService) GetChainMetrics(chainID isc.ChainID) *dto.ChainMetrics {
	chain := c.chainProvider().Get(chainID)
	if chain == nil {
		return nil
	}

	metrics := chain.GetNodeConnectionMetrics()
	registered := metrics.GetRegistered()

	return &dto.ChainMetrics{
		InAliasOutput:                   dto.MapMetricItem(metrics.GetInAliasOutput()),
		InOnLedgerRequest:               dto.MapMetricItem(metrics.GetInOnLedgerRequest()),
		InOutput:                        dto.MapMetricItem(metrics.GetInOutput()),
		InStateOutput:                   dto.MapMetricItem(metrics.GetInStateOutput()),
		InTxInclusionState:              dto.MapMetricItem(metrics.GetInTxInclusionState()),
		InMilestone:                     dto.MapMetricItem(metrics.GetInMilestone()),
		OutPublishGovernanceTransaction: dto.MapMetricItem(metrics.GetOutPublishGovernanceTransaction()),

		OutPullLatestOutput:          dto.MapMetricItem(metrics.GetOutPullLatestOutput()),
		OutPullOutputByID:            dto.MapMetricItem(metrics.GetOutPullOutputByID()),
		OutPullTxInclusionState:      dto.MapMetricItem(metrics.GetOutPullTxInclusionState()),
		OutPublisherStateTransaction: dto.MapMetricItem(metrics.GetOutPublishStateTransaction()),

		RegisteredChainIDs: registered,
	}
}

func (c *MetricsService) GetChainConsensusWorkflowMetrics(chainID isc.ChainID) *models.ConsensusWorkflowMetrics {
	chain := c.chainProvider().Get(chainID)
	if chain == nil {
		return nil
	}

	metrics := chain.GetConsensusWorkflowStatus()
	if metrics == nil {
		return nil
	}

	return models.MapConsensusWorkflowStatus(metrics)
}

func (c *MetricsService) GetChainConsensusPipeMetrics(chainID isc.ChainID) *models.ConsensusPipeMetrics {
	chain := c.chainProvider().Get(chainID)
	if chain == nil {
		return nil
	}

	metrics := chain.GetConsensusPipeMetrics()
	if metrics == nil {
		return nil
	}

	return models.MapConsensusPipeMetrics(metrics)
}
