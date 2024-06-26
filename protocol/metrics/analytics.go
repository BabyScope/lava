package metrics

import (
	"time"
)

type RelaySource int

const (
	SdkSource RelaySource = iota + 1
	GatewaySource
	BadgesSource
)

type RelayMetrics struct {
	ProjectHash  string
	Timestamp    time.Time
	ChainID      string
	APIType      string
	Latency      int64
	Success      bool
	ComputeUnits uint64
	Source       RelaySource
	Origin       string
	ApiMethod    string
}

type RelayAnalyticsDTO struct {
	ProjectHash  string
	Timestamp    string
	ChainID      string
	APIType      string
	Latency      uint64
	SuccessCount int64
	RelayCounts  int64
	TotalCu      uint64
	Source       RelaySource
	Origin       string
}

func NewRelayAnalytics(projectHash, chainId, apiType string) *RelayMetrics {
	return &RelayMetrics{
		Timestamp:   time.Now(),
		ProjectHash: projectHash,
		ChainID:     chainId,
		APIType:     apiType,
		Source:      GatewaySource,
	}
}
