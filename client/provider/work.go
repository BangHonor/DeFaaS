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
func (client *ProviderClient) Watcher(ctx context.Context, toBidderCh chan<- *data.DeploymentItem, errCh chan error) error {

	sink := make(chan *market.MarketNewDeploymentOrderEvent)
	sub, err := client.Market.Contract.MarketFilterer.WatchNewDeploymentOrderEvent(
		nil, sink, nil, nil, nil)
	if err != nil {
		log.Println("[provider/watcher] failed to watch event:", err)
		return err
	}

	go func() {

		for {
			select {

			case <-ctx.Done():
				return

			case err := <-sub.Err():
				log.Println("[provider/watcher] an error form subscription:", err)
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

				item := &data.DeploymentItem{}

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

func (client *ProviderClient) Bidder(ctx context.Context, fromWatcherCh <-chan *data.DeploymentItem, toPublisherCh chan<- *data.DeploymentItem, errCh chan error) error {

	go func() {

		for {
			select {

			case <-ctx.Done():
				return

			case _item := <-fromWatcherCh:

				//  start a new goroutine to do this time-heavy-cost task
				go func(item *data.DeploymentItem) {

					// change the state
					item.State = data.BiddingState

					item.Info.Provider = client.Key.Address
					item.Info.UnitPrice = item.Order.HighestUnitPrice

					// bid for this order
					txBid, err := client.Market.Bid(item.ID, item.Info.UnitPrice) // TODO smarter bidding strategy
					if err != nil {
						log.Printf("[provider/bidder] failed to bid for a deployment order, ID [%v]: %v\n", item.ID, err)
						errCh <- err
						return
					}
					if err := client.ComfirmTxByPolling(txBid.Hash(), basic.NumBlockToWaitRecommended); err != nil {
						log.Printf("[provider/bidder] failed to confirm tx[%v] for order, ID[%v]: %v\n", txBid.Hash(), item.ID, err)
						errCh <- err
						return
					}

					log.Printf("[provider/bidder] bid with unit price [%v] for a deployment order, ID [%v]\n", item.Info.UnitPrice, item.ID)

					// watch bidding end event
					sink := make(chan *market.MarketBiddingEndEvent)
					sub, err := client.Market.Contract.WatchBiddingEndEvent(nil, sink, []*big.Int{item.ID}, nil, nil)
					if err != nil {
						log.Println("[provider/bidder] failed to watch event:", err)
						errCh <- err
						return
					}

					// wait for bidding end event
					select {

					case <-ctx.Done():
						// this select done

					case err := <-sub.Err():
						log.Println("[provider/bidder] an error form subscription:", err)
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
				}(_item)

			}
		}
	}()

	return nil
}

func (client *ProviderClient) Fulfiller(ctx context.Context, fromBidderCh <-chan *data.DeploymentItem, errCh chan error) error {

	go func() {

		for {
			select {

			case <-ctx.Done():
				return

			case _item := <-fromBidderCh:

				//  start a new goroutine to do this time-heavy-cost task
				go func(item *data.DeploymentItem) {

					// publish
					{
						item.State = data.ConfirmingState

						item.Info.FuncPath = "funcPath"
						item.Info.DeployPath = GetDeployPathFromProviderConfig(client.providerConfig)

						// use customrt's public key to encrypt
						item.Info.AccessKey = "accessKey"
						encryptedAccessKey, err := data.EncryptAccessKey(item.Info.AccessKey, item.Order.PublicKey)
						if err != nil {
							log.Println("[provider/fulfiller] failed to encrypted accessKey")
							errCh <- err
							return
						}

						// publish
						txPublish, err := client.Market.PublishDeploymentInfo(item.ID, item.Info.FuncPath, item.Info.DeployPath, encryptedAccessKey)
						if err != nil {
							log.Printf("[provider/fulfiller] failed to publish deployment information for a deployment order, ID [%v]: %v\n", item.ID, err)
							errCh <- err
							return
						}
						if err := client.ComfirmTxByPolling(txPublish.Hash(), basic.NumBlockToWaitRecommended); err != nil {
							log.Printf("[provider/fulfiller] failed to confirm tx[%v] for order, ID[%v]: %v\n", txPublish.Hash(), item.ID, err)
							errCh <- err
							return
						}

						log.Printf("[provider/fulfiller] publish a new deployment info for the order, ID [%v]\n", item.ID)

						// watch new lease event
						sink := make(chan *market.MarketNewLeaseEvent)
						sub, err := client.Market.Contract.WatchNewLeaseEvent(nil, sink,
							[]*big.Int{item.ID},
							[]common.Address{item.Order.Customer},
							[]common.Address{item.Info.Provider})
						if err != nil {
							log.Println("[provider/fulfiller] failed to watch event:", err)
							errCh <- err
							return
						}

						// wait for new lease event
						select {

						case <-ctx.Done():
							// this select done

						case err := <-sub.Err():
							log.Println("[provider/fulfiller] an error form subscription:", err)
							errCh <- err

						case event := <-sink:

							// -------------- event --------------------
							// DeploymentOrderID *big.Int
							// Customer          common.Address
							// Provider          common.Address

							if helper.EqualAddress(item.Order.Customer, event.Customer) && helper.EqualAddress(item.Info.Provider, event.Provider) {

								log.Printf("[provider/publisher] new lease, ID [%v]\n", item.ID)
							}
						}

					}

					// deploy & fulfill
					{
						item.State = data.DeployingState

						// register to deploy server
						// and wait until deploy server notifys back
						// item is read-only in `client.DeployServerRegisterWait` !!!
						log.Printf("[provider/fulfiller] wait for deployment action from customer, ID [%v] ...\n", item.ID)
						item.Order.FulfillSecretKey = client.DeployServerRegisterWait(item.ID)

						txFulfill, err := client.Market.FulfillDeploymentOrder(item.ID, item.Order.FulfillSecretKey)
						if err != nil {
							log.Printf("[provider/fulfiller] failed to fulfill for a deployment order, ID [%v]: %v\n", item.ID, err)
							errCh <- err
							return
						}
						if err := client.ComfirmTxByPolling(txFulfill.Hash(), basic.NumBlockToWaitRecommended); err != nil {
							log.Printf("[provider/fulfiller] failed to confirm tx[%v] for order, ID[%v]: %v\n", txFulfill.Hash(), item.ID, err)
							errCh <- err
							return
						}

						item.State = data.FulfillingState

						log.Printf("[provider/fulfiller] fulfilling deployment order, ID[%v]\n", item.ID)
					}

					// finish
					{
						// watch finish event
						// item.State = data.FinishedState
						// TODO
					}

				}(_item)

			}
		}
	}()

	return nil
}
