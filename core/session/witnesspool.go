package session

import (
	"context"
	"defaas/contracts/go/witnesspool"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// NewWitnessPoolSeesion ...
func NewWitnessPoolSeesion(backend bind.ContractBackend, address common.Address, auth *bind.TransactOpts) (*witnesspool.WitnessPoolSession, error) {

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
