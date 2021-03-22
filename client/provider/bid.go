package provider

import (
	"context"
	"defaas/client/basic"
	"defaas/contracts/go/market"
	"defaas/core/helper"
	"log"
	"math/big"
	"sync"
)

// BidFilter is afilter rule for bidding new deployment order.
type BidFilter struct {
	LongestFaaSDuration *big.Int
	LowestUnitPrice     *big.Int

	// TODO more filter field
}

// Check checks whether the new deployment order is legal.
func (f *BidFilter) Check(event *market.MarketNewDeploymentOrderEvent) bool {

	// TODO

	return true
}

// Filter execute custom filter rules.
func (f *BidFilter) Filter(event *market.MarketNewDeploymentOrderEvent) bool {

	if f.LowestUnitPrice != nil && f.LowestUnitPrice.Cmp(event.HighestUnitPrice) == 1 {
		return false
	}

	// TODO more filter rule

	return true
}

// ------------------------------------------------------------------------------------------------

// BidStrategy is a bidding strategy, including bids for different FaaS levels.
type BidStrategy struct {
	// TODO
}

// ------------------------------------------------------------------------------------------------

type Bidder struct {
	sync.Mutex
	Err      chan error
	Filter   *BidFilter
	Strategy *BidStrategy
}

func (client *ProviderClient) Bidding(ctx context.Context) error {

	filter := &BidFilter{}
	strategy := &BidStrategy{}                                                  // TODO support as a parameter
	bidder, err := client.NewBidder(ctx, client.sinkBiddenID, filter, strategy) // start a daemons for bidding

	if err != nil {
		return err
	}

	// start a gorountine
	// customer of of bidden deployment order
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-bidder.Err:
				// just print err
				log.Println(err)
			case id := <-client.sinkBiddenID:
				go client.Publish(ctx, id)
			}
		}
	}()

	return nil
}

func (client *ProviderClient) NewBidder(ctx context.Context, sinkBiddenID chan<- *big.Int, filter *BidFilter, strategy *BidStrategy) (*Bidder, error) {

	subBid := &Bidder{
		Err:      make(chan error),
		Filter:   filter,
		Strategy: strategy,
	}

	// construct subscription of event `new deployemnt order`
	sinkNewOrder := make(chan *market.MarketNewDeploymentOrderEvent)
	subNewOrder, err := client.Market.Contract.MarketFilterer.WatchNewDeploymentOrderEvent(
		nil, sinkNewOrder, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	bidFor := func(ctx context.Context, id *big.Int, unitPirce *big.Int) {

		// bid for
		txBid, err := client.Market.Bid(id, unitPirce)
		if err != nil {
			subBid.Err <- err
		}
		if err := client.ComfirmTxByPolling(txBid.Hash(), basic.NumBlockToWaitRecommended); err != nil {
			subBid.Err <- err
		}

		// watch bidding end event
		sinkBiddingEnd := make(chan *market.MarketBiddingEndEvent)
		subBiddingEnd, err := client.Market.Contract.WatchBiddingEndEvent(nil, sinkBiddingEnd, []*big.Int{id}, nil, nil)
		if err != nil {
			subBid.Err <- err
			return
		}

		select {

		case <-ctx.Done():
			return

		case err := <-subBiddingEnd.Err():
			subBid.Err <- err
			return

		case event := <-sinkBiddingEnd:

			if helper.EqualAddress(client.Key.Address, event.Provider) {
				// bidding successfully
				sinkBiddenID <- id
			}
			return
		}

	}

	// watch new order
	// producer of bidden deployment order
	go func() {
		for {
			select {

			case <-ctx.Done():
				return

			case err := <-subNewOrder.Err():
				subBid.Err <- err
				return

			case event := <-sinkNewOrder:

				if subBid.Filter.Check(event) && subBid.Filter.Filter(event) {
					go bidFor(
						ctx, event.DeploymentOrderID, event.HighestUnitPrice)
				}

			}
		}
	}()

	return subBid, nil
}
