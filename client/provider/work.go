package provider

import (
	"context"
	"defaas/client/basic"
	"defaas/contracts/go/market"
	"defaas/core/data"
	"defaas/core/helper"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Watcher watches `NewDeploymentOrderEvent` from `Market` contract
func (client *ProviderClient) Watcher(ctx context.Context, toBidderCh chan<- *DeploymentItem, errCh chan error) error {

	sink := make(chan *market.MarketNewDeploymentOrderEvent)
	sub, err := client.Market.Contract.MarketFilterer.WatchNewDeploymentOrderEvent(
		nil, sink, nil, nil, nil)

	if err != nil {
		return err
	}

	go func() {

		for {
			select {

			case <-ctx.Done():
				return

			case err := <-sub.Err():
				errCh <- err
				return

			case event := <-sink:

				// -------------- event --------------------
				// Customer          common.Address
				// Nonce             *big.Int
				// DeploymentOrderID *big.Int
				// FaasLevelID       *big.Int
				// HighestUnitPrice  *big.Int
				// FaasDuration      *big.Int
				// BiddingDuration   *big.Int
				// PublicKey         string

				// -------------- order --------------------
				// Customer         common.Address
				// Nonce            *big.Int
				// FaaSLevelID      *big.Int
				// HighestUnitPrice *big.Int
				// FaaSDuration     *big.Int
				// BiddingDuration  *big.Int
				// PublicKey        string
				// FulfillSecretKey [32]byte
				// FulfillKey       [32]byte

				item := &DeploymentItem{}

				item.State = data.InitState
				item.ID = event.DeploymentOrderID

				item.Order.Customer = event.Customer
				item.Order.Nonce = event.Nonce
				item.Order.FaaSLevelID = event.FaasLevelID
				item.Order.HighestUnitPrice = event.HighestUnitPrice
				item.Order.FaaSDuration = event.FaasDuration
				item.Order.BiddingDuration = event.BiddingDuration
				item.Order.PublicKey = event.PublicKey

				log.Printf("[provider/watcher] watch a new deployment order, ID [%v]\n", item.ID)

				client.itemPool.Set(item.ID, item)

				toBidderCh <- item
			}
		}
	}()

	return nil
}

func (client *ProviderClient) Bidder(ctx context.Context, fromWatcherCh <-chan *DeploymentItem, toPublisherCh chan<- *DeploymentItem, errCh chan error) error {

	go func() {

		for {
			select {

			case <-ctx.Done():
				return

			case item := <-fromWatcherCh:

				// change the state
				item.State = data.BiddingState

				item.Info.Provider = client.Key.Address
				item.Info.UnitPrice = item.Order.HighestUnitPrice

				// bid for this order
				txBid, err := client.Market.Bid(item.ID, item.Info.UnitPrice) // TODO smarter bidding strategy
				if err != nil {
					errCh <- err
					continue
				}
				if err := client.ComfirmTxByPolling(txBid.Hash(), basic.NumBlockToWaitRecommended); err != nil {
					errCh <- err
					continue
				}

				// watch bidding end event
				sink := make(chan *market.MarketBiddingEndEvent)
				sub, err := client.Market.Contract.WatchBiddingEndEvent(nil, sink, []*big.Int{item.ID}, nil, nil)
				if err != nil {
					errCh <- err
					continue
				}

				// wait for bidding end event
				select {

				case <-ctx.Done():
					// this select done

				case err := <-sub.Err():
					errCh <- err

				case event := <-sink:

					// -------------- event --------------------
					// DeploymentOrderID *big.Int
					// Provider          common.Address
					// Success           bool
					// FaasLevelID       *big.Int
					// UnitPrice         *big.Int

					// check whether I win the bidding
					if event.Success && helper.EqualAddress(item.Info.Provider, event.Provider) {

						// I win the bidding
						// send the bidden item to next processing stage
						toPublisherCh <- item

						log.Printf("[provider/bidder] win the bidding of a deployment order, ID [%v]\n", item.ID)

					} else {

						client.itemPool.Remove(item.ID)

						log.Printf("[provider/bidder] failed to bid for a deployment order, ID [%v]\n", item.ID)
					}
				}

			}
		}
	}()

	return nil
}

func (client *ProviderClient) Publisher(ctx context.Context, fromBidderCh <-chan *DeploymentItem, toFulfillerCh chan<- *DeploymentItem, errCh chan error) error {

	go func() {

		for {
			select {

			case <-ctx.Done():
				return

			case item := <-fromBidderCh:

				item.State = data.ConfirmingState

				item.Info.FuncPath = "funcPath"
				item.Info.DeployPath = client.providerConfig.DeployPath
				item.Info.AccessKey = "accessKey"

				// publish
				txPublish, err := client.Market.PublishDeploymentInfo(item.ID, item.Info.FuncPath, item.Info.DeployPath, item.Info.AccessKey)
				if err != nil {
					errCh <- err
					continue
				}
				if err := client.ComfirmTxByPolling(txPublish.Hash(), basic.NumBlockToWaitRecommended); err != nil {
					errCh <- err
					continue
				}

				// watch new lease event
				sink := make(chan *market.MarketNewLeaseEvent)
				sub, err := client.Market.Contract.WatchNewLeaseEvent(nil, sink,
					[]*big.Int{item.ID},
					[]common.Address{item.Order.Customer},
					[]common.Address{item.Info.Provider})
				if err != nil {
					errCh <- err
					continue
				}

				// wait for new lease event
				select {

				case <-ctx.Done():
					// this select done

				case err := <-sub.Err():
					errCh <- err

				case event := <-sink:

					// -------------- event --------------------
					// DeploymentOrderID *big.Int
					// Customer          common.Address
					// Provider          common.Address

					if helper.EqualAddress(item.Order.Customer, event.Customer) && helper.EqualAddress(item.Info.Provider, event.Provider) {

						log.Printf("[provider/publisher] new lease, ID [%v]\n", item.ID)
						toFulfillerCh <- item

					}
				}
			}
		}
	}()

	return nil
}

func (client *ProviderClient) Fulfiller(ctx context.Context, inCh <-chan *DeploymentItem, outCh chan<- *DeploymentItem, errCh chan error) error {
	return nil
}

func (client *ProviderClient) Finisher(ctx context.Context, inCh <-chan *DeploymentItem, errCh chan error) error {
	return nil
}
