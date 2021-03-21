package customer

// https://github.com/dyrkin/fsm

const (
	InitState       = "Init"
	BiddingState    = "Bidding"
	ConfirmingState = "Confirming"
	DeployingState  = "Deploying"
	FulfillingState = "Fulfilling"
	SettledState    = "Settled"
	FinishedState   = "Finished"
)

// func (client *CustomerClient) Deploy(order *data.DeploymentOrder, task interface{}) error {

// 	// 1 new
// 	// watch new deployemnt order event (get id)
// 	// wait for bidding
// 	// 3 match
// 	// 3 query match result (get match result)
// 	// watch new deployment info event (get info)
// 	// 4b confirm
// 	// watch new lease event
// 	// 5a off-chain deploying
// 	// 5b fulfill
// 	// 6 finish

// 	newTx, err := client.Market.NewDeploymentOrder(
// 		order.FaaSLevelID,
// 		order.HighestUnitPrice,
// 		order.FaaSDuration,
// 		order.BiddingDuration)

// 	if err != nil {
// 		return nil
// 	}
// 	if err := client.ComfirmTxByPolling(newTx.Hash(), basic.NumBlockToWaitRecommended); err != nil {
// 		return err
// 	}

// 	newOrderSub, err := client.Market.Contract.MarketFilterer.WatchNewDeploymentOrderEvent()

// 	// new
// 	cr.Send(order)

// 	// wait for bidding
// 	time.Sleep(time.Second)

// 	// match
// 	cr.Send("match")

// 	// confirm
// 	confirmTimeout := time.Timer(time.Second)

// ConfirmLoop:
// 	for {
// 		select {
// 		case confirmTimeout.C:
// 			cr.Send("timeout")
// 			return errors.New("confirm timeout")
// 			// case sub <-
// 			// cr.Send(deployInfo)
// 			// break ConfirmLoop
// 		}
// 	}

// 	// deploy

// 	// 部署到此为止

// 	// fulfill
// 	// time.Sleep(time.Hour)

// 	return nil
// }

// func (client *CustomerClient) finish(id *big.Int) {
// 	// TODO
// }
