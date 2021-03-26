package session

import (
	"context"
	"math/big"
	"testing"

	"defaas/contracts/go/faastoken"

	"github.com/stretchr/testify/assert"
)

func TestFaaSTokenSessionDeploy(t *testing.T) {

	assert := assert.New(t)

	auth, blockchain := newSimBlockchain()

	_, tx, _, err := faastoken.DeployFaaSToken(auth, blockchain)
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
}

func TestFaaSTokenMind(t *testing.T) {

	assert := assert.New(t)

	auth, blockchain, _, session := newSimFaaSToken()

	before, _ := session.TotalSupply()

	mindedAmount := big.NewInt(100)
	_, _ = session.MintTo(auth.From, mindedAmount)
	blockchain.Commit()

	after, _ := session.TotalSupply()

	expected := big.NewInt(0).Add(before, mindedAmount)
	assert.Equal(0, after.Cmp(expected))
}

func TestFaaSTokenBalanceOf(t *testing.T) {

	assert := assert.New(t)

	auth, blockchain, _, session := newSimFaaSToken()

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
