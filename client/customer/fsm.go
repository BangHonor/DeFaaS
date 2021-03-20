package customer

// https://github.com/dyrkin/fsm

import (
	"defaas/core/data"
	"errors"
	"math/big"
	"time"
)

const (
	InitState       = "Init"
	BiddingState    = "Bidding"
	ConfirmingState = "Confirming"
	DeployingState  = "Deploying"
	FulfillingState = "Fulfilling"
	SettledState    = "Settled"
	FinishedState   = "Finished"
)

func (client *CustomerClient) Deploy(order *data.DeploymentOrder, task interface{}) error {

	cr := newCustomerRun()

	// new
	cr.Send(order)

	// wait for bidding
	time.Sleep(time.Second)

	// match
	cr.Send("match")

	// confirm
	confirmTimeout := time.Timer(time.Second)

ConfirmLoop:
	for {
		select {
		case confirmTimeout.C:
			cr.Send("timeout")
			return errors.New("confirm timeout")
			// case sub <-
			// cr.Send(deployInfo)
			// break ConfirmLoop
		}
	}

	// deploy

	// 部署到此为止

	// fulfill
	// time.Sleep(time.Hour)

	// newTx, err := client.Market.NewDeploymentOrder(
	// 	order.FaaSLevelID,
	// 	order.HighestUnitPrice,
	// 	order.FaaSDuration,
	// 	order.BiddingDuration)

	// if err != nil {
	// 	return nil
	// }

	// if err := client.ComfirmTxByPolling(newTx.Hash(), basic.NumBlockToWaitRecommended); err != nil {
	// 	return err
	// }

	// go client.deployRun()

	return nil
}

func (client *CustomerClient) finish(id *big.Int) {
	// TODO
}
