package suite

import (
	"context"
	"defaas/contracts/go/faastoken"
	"defaas/contracts/go/market"
	"defaas/contracts/go/witnesspool"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type FaaSTokenSession = faastoken.FaaSTokenSession
type MarketSeesion = market.MarketSession
type WitnessPoolSession = witnesspool.WitnessPoolSession

type Suite struct {
	FaaSToken   *FaaSTokenSession
	Market      *MarketSeesion
	WitnessPool *WitnessPoolSession
}

func NewSuite(
	backend bind.ContractBackend, auth *bind.TransactOpts,
	faastokenAddress, marketAddress, witnesspoolAddress common.Address) (*Suite, error) {

	var (
		err error
	)

	suite := &Suite{}

	suite.FaaSToken, err = NewFaaSTokenSeesion(backend, faastokenAddress, auth)
	if err != nil {
		return nil, err
	}

	suite.Market, err = NewMarketSeesion(backend, marketAddress, auth)
	if err != nil {
		return nil, err
	}

	suite.WitnessPool, err = NewWitnessPoolSession(backend, witnesspoolAddress, auth)
	if err != nil {
		return nil, err
	}

	return suite, nil
}

// NewFaaSTokenSeesion ...
func NewFaaSTokenSeesion(backend bind.ContractBackend, address common.Address, auth *bind.TransactOpts) (*FaaSTokenSession, error) {

	instance, err := faastoken.NewFaaSToken(address, backend)
	if err != nil {
		return nil, err
	}

	session := &faastoken.FaaSTokenSession{
		Contract: instance,
		CallOpts: bind.CallOpts{
			// Pending: true,
			Pending: false,
		},
		TransactOpts: bind.TransactOpts{
			From:    auth.From,
			Signer:  auth.Signer,
			Context: context.TODO(),
		},
	}

	return session, nil
}

func NewMarketSeesion(backend bind.ContractBackend, address common.Address, auth *bind.TransactOpts) (*market.MarketSession, error) {

	instance, err := market.NewMarket(address, backend)
	if err != nil {
		return nil, err
	}

	session := &market.MarketSession{
		Contract: instance,
		CallOpts: bind.CallOpts{
			// Pending: true,
			Pending: false,
		},
		TransactOpts: bind.TransactOpts{
			From:    auth.From,
			Signer:  auth.Signer,
			Context: context.TODO(),
		},
	}

	return session, nil
}

// NewWitnessPoolSession ...
func NewWitnessPoolSession(backend bind.ContractBackend, address common.Address, auth *bind.TransactOpts) (*witnesspool.WitnessPoolSession, error) {

	instance, err := witnesspool.NewWitnessPool(address, backend)
	if err != nil {
		return nil, err
	}

	session := &witnesspool.WitnessPoolSession{
		Contract: instance,
		CallOpts: bind.CallOpts{
			// Pending: true,
			Pending: false,
		},
		TransactOpts: bind.TransactOpts{
			From:    auth.From,
			Signer:  auth.Signer,
			Context: context.TODO(),
		},
	}

	return session, nil
}
