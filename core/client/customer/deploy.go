package customer

import (
	market "defaas/contracts/go/market"
	"defaas/core/client/basic"
	"defaas/core/helper"
	data "defaas/core/model"
	"errors"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

func (client *CustomerClient) NewDeploy(faasLevelID, faasDuration, highestUnitPrice, biddingDuration *big.Int) error {

	order := &data.DeploymentOrder{
		Customer:         client.Key.Address,
		Nonce:            big.NewInt(atomic.AddInt64(&client.nonce, 1)),
		FaaSLevelID:      faasLevelID,
		HighestUnitPrice: highestUnitPrice,
		FaaSDuration:     faasDuration,
		BiddingDuration:  biddingDuration,
	}

	if order.BiddingDuration == nil {
		order.BiddingDuration = BiddingDurationRecommended
	}

	err := client.DeployOrder(order)
	if err != nil {
		return err
	}

	return nil
}

func (client *CustomerClient) DeployOrder(_order *data.DeploymentOrder) error {

	// overall
	// --------------------------------------------------------
	// 1 [new]
	// new
	// watch new deployemnt order event (get id)
	// 2 [match]
	// wait for bidding
	// match
	// query match result (get match result)
	// 3 [confirm]
	// watch new deployment info event (get info)
	// confirm
	// query lease (get confirm result)
	// 4 [deploy]
	// off-chain deploying
	// watch new SLA event
	// 5 [finish]
	// finish
	// --------------------------------------------------------

	var (
		err    error                 = nil
		item   *data.DeploymentItem  = new(data.DeploymentItem)
		filter market.MarketFilterer = client.Market.Contract.MarketFilterer

		// events to watch
		subNewOrder event.Subscription = nil
		subNewInfo  event.Subscription = nil
		subNewSLA   event.Subscription = nil
	)

	if !helper.EqualAddress(_order.Customer, client.Key.Address) {
		err = errors.New("wrong customer address in structure 'order'")
		return err
	}

	// 1 [new]
	{
		item.State = data.InitState
		item.Order = *_order

		if err = client._new(item); err != nil {
			return err
		}

		// watch new deployemnt order event
		sinkNewOrder := make(chan *market.MarketNewDeploymentOrderEvent)
		subNewOrder, err = filter.WatchNewDeploymentOrderEvent(nil, sinkNewOrder,
			[]common.Address{item.Order.Customer},
			[]*big.Int{item.Order.Nonce},
			nil)
		if err != nil {
			return err
		}

		timeoutNewOrder := time.NewTimer(1 * time.Minute)

		select {
		case <-timeoutNewOrder.C:
			err = errors.New("time out while waiting for new deployment order event")
			return err
		case err = <-subNewOrder.Err():
			return err
		case event := <-sinkNewOrder:
			item.ID = event.DeploymentOrderID // set ID
		}
	}

	client.itemPool.Set(item.ID, item) // add new item to pool
	defer func() {
		if err != nil {
			client.itemPool.Remove(item.ID) // if failed, remove the item
		}
	}()

	// 2 [match]
	{
		item.State = data.BiddingState

		// wait for bidding
		// 1 min is a redundancy
		time.Sleep(
			(1 * time.Minute) +
				(time.Duration(item.Order.BiddingDuration.Int64()) * time.Second))

		// match
		if err = client._match(item); err != nil {
			return nil
		}

		// query match result (get match result)
		var isMatch bool
		isMatch, err = client.Market.QueryMatch(item.ID)

		if err != nil {
			return err
		}
		if !isMatch {
			err = errors.New("deployment order not matched")
			return err
		}
	}

	// 3 [confirm]
	{
		item.State = data.ConfirmingState

		// watch new deployment info event
		sinkNewInfo := make(chan *market.MarketNewDeploymentInfoEvent)
		subNewInfo, err = filter.WatchNewDeploymentInfoEvent(
			nil, sinkNewInfo, []*big.Int{item.ID}, nil)
		if err != nil {
			return err
		}

		timeoutNewInfo := time.NewTimer(1 * time.Minute)

		select {
		case err = <-subNewInfo.Err():
			return err
		case <-timeoutNewInfo.C:
			err = errors.New("time out while waiting for new deployment information event")
			return err
		case event := <-sinkNewInfo:
			item.Info.Provider = event.Provider
			item.Info.FuncPath = event.FuncPath
			item.Info.DeployPath = event.DeployPath

		}

		// confirm
		if err = client._confirm(item); err != nil {
			return err
		}

		// query lease
		var customer common.Address
		var provider common.Address
		customer, provider, err = client.Market.QueryLease(item.ID)
		if err != nil {
			return err
		}
		if customer != item.Order.Customer || provider != item.Info.Provider {
			err = errors.New("fail to create lease")
			return err
		}
	}

	// 4 [deploy]
	{
		item.State = data.DeployingState

		// 5a off-chain deploying
		// ???
		// TODO

		// watch new SLA event
		sinkNewSLA := make(chan *market.MarketNewSLAEvent)
		subNewSLA, err = filter.WatchNewSLAEvent(nil, sinkNewSLA,
			[]*big.Int{item.ID})
		if err != nil {
			return err
		}
		timeoutNewSLA := time.NewTimer(1 * time.Minute)

		select {
		case err = <-subNewSLA.Err():
			return err
		case event := <-sinkNewSLA:
			_ = event
		case <-timeoutNewSLA.C:
			err = errors.New("time out while waiting for new SLA event")
			return err
		}

		item.State = data.FulfillingState
	}

	// 5 [finish]
	{
		if err = client._finish(item); err != nil {
			return err
		}

		item.State = data.FinishedState
	}

	return nil
}

// -----------------------------------------------------------------------------------------------------------------

func (client *CustomerClient) _new(item *data.DeploymentItem) error {

	if item.State != data.InitState {
		return data.ErrWrongState
	}

	txNewOrder, err := client.Market.NewDeploymentOrder(
		item.Order.Nonce,
		item.Order.FaaSLevelID,
		item.Order.FaaSDuration,
		"", // funcmsg TOOD
		item.Order.HighestUnitPrice,
		item.Order.BiddingDuration)

	if err != nil {
		return err
	}

	if err := client.ComfirmTxByPolling(txNewOrder.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	return nil
}

func (client *CustomerClient) _match(item *data.DeploymentItem) error {

	if item.State != data.BiddingState {
		return data.ErrWrongState
	}

	txMatchOrder, err := client.Market.MatchDeploymentOrder(item.ID)

	if err != nil {
		return err
	}

	if err := client.ComfirmTxByPolling(txMatchOrder.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	return nil
}

func (client *CustomerClient) _confirm(item *data.DeploymentItem) error {

	if item.State != data.ConfirmingState {
		return data.ErrWrongState
	}

	txConfirm, err := client.Market.ConfirmDeploymentInfo(item.ID)

	if err != nil {
		return err
	}

	if err := client.ComfirmTxByPolling(txConfirm.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	return nil
}

func (client *CustomerClient) _finish(item *data.DeploymentItem) error {

	if item.State != data.FulfillingState {
		return data.ErrWrongState
	}

	txFinish, err := client.Market.FinishDeploymemtOrder(item.ID)

	if err != nil {
		return nil
	}

	if err := client.ComfirmTxByPolling(txFinish.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	return nil
}
