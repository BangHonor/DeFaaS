package provider

import (
	"defaas/contracts/go/market"
	"math/big"
)

// BidFilter is afilter rule for bidding new deployment order.
type BidFilter struct {
	LongestFaaSDuration *big.Int
	LowestUnitPrice     *big.Int
	// TODO
}

// Check checks whether the new deployment order is legal.
func (f *BidFilter) Check(event *market.MarketNewDeploymentOrderEvent) bool {
	return true
}

// Filter execute custom filter rules.
func (f *BidFilter) Filter(event *market.MarketNewDeploymentOrderEvent) bool {

	if f.LowestUnitPrice != nil && f.LowestUnitPrice.Cmp(event.HighestUnitPrice) == 1 {
		return false
	}

	return true
}

// ------------------------------------------------------------------------------------------------

// BidStrategy is a bidding strategy, including bids for different FaaS levels.
type BidStrategy struct {
	// TODO
}

// ------------------------------------------------------------------------------------------------
