package session

import (
	"context"
	"defaas/contracts/go/faastoken"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type FaaSTokenSession = faastoken.FaaSTokenSession

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
