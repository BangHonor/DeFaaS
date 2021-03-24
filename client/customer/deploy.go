package customer

import (
	basic "defaas/client/basic"
	market "defaas/contracts/go/market"
	"defaas/core/data"
	"defaas/core/helper"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func (client *CustomerClient) Deploy(order *data.DeploymentOrder, task interface{}) error {

	// overall
	// --------------------------------------------------------
	// 1 new
	// watch new deployemnt order event (get id)
	// wait for bidding
	// 3 match
	// 3 query match result (get match result)
	// watch new deployment info event (get info)
	// 4b confirm
	// 4b query lease (get confirm result)
	// 5a off-chain deploying
	// watch new SLA event
	// 6 finish

	if !helper.EqualAddress(order.Customer, client.Key.Address) {
		return errors.New("wrong customer address in structure 'order'")
	}

	var (
		id     *big.Int // deployment id
		info   data.DeploymentInfo
		filter market.MarketFilterer = client.Market.Contract.MarketFilterer
	)

	// 1 new
	txNewOrder, err := client.Market.NewDeploymentOrder(
		order.Nonce,
		order.FaaSLevelID,
		order.HighestUnitPrice,
		order.FaaSDuration,
		order.BiddingDuration,
		order.PublicKey)

	if err != nil {
		return nil
	}
	if err := client.ComfirmTxByPolling(txNewOrder.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	// watch new deployemnt order event
	sinkNewOrder := make(chan *market.MarketNewDeploymentOrderEvent)
	subNewOrder, err := filter.WatchNewDeploymentOrderEvent(
		nil, sinkNewOrder, []common.Address{order.Customer}, []*big.Int{order.Nonce}, nil)
	if err != nil {
		return err
	}
	timeoutNewOrder := time.NewTimer(1 * time.Minute)

	select {
	case err := <-subNewOrder.Err():
		return err
	case event := <-sinkNewOrder:
		id = event.DeploymentOrderID
	case <-timeoutNewOrder.C:
		return errors.New("time out while waiting for new delpoyment order event")
	}

	// wait for bidding
	// 1 min is a redundancy
	time.Sleep(
		(1 * time.Minute) +
			(time.Duration(order.BiddingDuration.Int64()) * time.Second))

	// 3 match
	txMatchOrder, err := client.Market.MatchDeploymentOrder(id)

	if err != nil {
		return nil
	}
	if err := client.ComfirmTxByPolling(txMatchOrder.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	// 3 query match result (get match result)
	isMatch, err := client.Market.QueryMatch(id)
	if err != nil {
		return err
	}
	if !isMatch {
		return errors.New("deployment order not matched")
	}

	// watch new deployment info event
	sinkNewInfo := make(chan *market.MarketNewDeploymentInfoEvent)
	subNewInfo, err := filter.WatchNewDeploymentInfoEvent(
		nil, sinkNewInfo, []*big.Int{id}, nil)
	if err != nil {
		return err
	}
	timeoutNewInfo := time.NewTimer(1 * time.Minute)

	select {
	case err := <-subNewInfo.Err():
		return err
	case event := <-sinkNewInfo:
		info.Provider = event.Provider
		info.FuncPath = event.FuncPath
		info.DeployPath = event.DeployPath
		info.AccessKey = event.AccessKey
	case <-timeoutNewInfo.C:
		return errors.New("time out while waiting for new delpoyment infomation event")
	}

	// TODO check `info.DeployPath` before confirm

	// 4b confirm
	txConfirm, err := client.Market.ConfirmDeploymentInfo(id, order.FulfillKey)

	if err != nil {
		return nil
	}
	if err := client.ComfirmTxByPolling(txConfirm.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	// query lease (get confirm result)
	customer, provider, err := client.Market.QueryLease(id)
	if err != nil {
		return err
	}
	if customer != order.Customer || provider != info.Provider {
		return errors.New("fail to create lease")
	}

	// 5a off-chain deploying
	// TODO
	// Adapter
	// send `order.FulfillSecretKey` to provider

	// watch new SLA event
	sinkNewSLA := make(chan *market.MarketNewSLAEvent)
	subNewSLA, err := filter.WatchNewSLAEvent(
		nil, sinkNewSLA, []*big.Int{id})
	if err != nil {
		return err
	}
	timeoutNewSLA := time.NewTimer(1 * time.Minute)

	select {
	case err := <-subNewSLA.Err():
		return err
	case event := <-sinkNewSLA:
		_ = event
	case <-timeoutNewSLA.C:
		return errors.New("time out while waiting for new SLA event")
	}

	// 6 finish
	if err := client.finish(id); err != nil {
		return err
	}

	return nil
}

func (client *CustomerClient) finish(id *big.Int) error {

	txFinish, err := client.Market.FinishDeploymemtOrder(id)

	if err != nil {
		return nil
	}
	if err := client.ComfirmTxByPolling(txFinish.Hash(), basic.NumBlockToWaitRecommended); err != nil {
		return err
	}

	return nil
}
