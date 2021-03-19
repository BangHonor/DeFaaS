package session

import (
	"context"
	"math/big"
	"testing"

	"defaas/contracts/go/faastoken"
	"defaas/core/helper"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestFaaSTokenSessionDeploy(t *testing.T) {

	assert := assert.New(t)

	auth, blockchain := helper.NewSim()

	address, tx, _, err := faastoken.DeployFaaSToken(auth, blockchain)
	if err != nil {
		t.Fatal(err)
	}

	// test pedding
	{
		if _, isPedding, err := blockchain.TransactionByHash(context.TODO(), tx.Hash()); err != nil {
			t.Fatal(err)
		} else {
			assert.True(isPedding)
		}

		blockchain.Commit()

		if _, isPedding, err := blockchain.TransactionByHash(context.TODO(), tx.Hash()); err != nil {
			t.Fatal(err)
		} else if isPedding == false {
			assert.False(isPedding)
		}
	}

	session, err := NewFaaSTokenSeesion(blockchain, address, auth)
	if err != nil {
		t.Fatal(err)
	}

	var (
		totolSupply  *big.Int
		mindedAmount *big.Int = big.NewInt(100)
		balanceOf    *big.Int
	)

	totolSupply, _ = session.TotalSupply()
	assert.Equal(0, totolSupply.Cmp(big.NewInt(0)))

	balanceOf, _ = session.BalanceOf(auth.From)
	assert.Equal(0, balanceOf.Cmp(big.NewInt(0)))

	_, _ = session.Mint(mindedAmount)
	blockchain.Commit()

	totolSupply, _ = session.TotalSupply()
	assert.Equal(0, totolSupply.Cmp(mindedAmount))

	balanceOf, _ = session.BalanceOf(auth.From)
	assert.Equal(0, balanceOf.Cmp(mindedAmount))
}

func getSimFaaSToken() (*bind.TransactOpts, *backends.SimulatedBackend, common.Address, *faastoken.FaaSTokenSession) {

	auth, blockchain := helper.NewSim()

	address, _, _, _ := faastoken.DeployFaaSToken(auth, blockchain)
	blockchain.Commit()
	session, _ := NewFaaSTokenSeesion(blockchain, address, auth)

	return auth, blockchain, address, session
}

func TestFaaSTokenMind(t *testing.T) {

	assert := assert.New(t)

	auth, blockchain, _, session := getSimFaaSToken()

	before, _ := session.TotalSupply()

	mindedAmount := big.NewInt(100)
	_, _ = session.MintTo(auth.From, mindedAmount)
	blockchain.Commit()

	after, _ := session.TotalSupply()

	expected := big.NewInt(0).Add(before, mindedAmount)
	assert.Equal(0, after.Cmp(expected))
}
