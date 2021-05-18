package helper

import (
	"context"
	"defaas/contracts/go/faastoken"
	"log"
)

func FaaSTokenEventSub(ctx context.Context, contract *faastoken.FaaSToken) error {

	var (
		filter       = contract.FaaSTokenFilterer
		sinkApproval = make(chan *faastoken.FaaSTokenApproval)
		sinkTransfer = make(chan *faastoken.FaaSTokenTransfer)
	)

	subApproval, err := filter.WatchApproval(nil, sinkApproval, nil, nil)
	if err != nil {
		return err
	}

	subTransfer, err := filter.WatchTransfer(nil, sinkTransfer, nil, nil)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-subApproval.Err():
				log.Fatal(err)
				return
			case event := <-sinkApproval:
				log.Printf("[event] [Approval] [%v]--[%v]-->[%v]\n", event.Owner, event.Value, event.Spender)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-subTransfer.Err():
				log.Fatal(err)
				return
			case event := <-sinkTransfer:
				log.Printf("[event] [Transfer] [%v]--[%v]-->[%v]\n", event.From, event.Value, event.To)
			}
		}
	}()

	return nil
}
