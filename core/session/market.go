package session

import (
	"context"
	"defaas/contracts/go/market"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// NewMarketSeesion ...
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
