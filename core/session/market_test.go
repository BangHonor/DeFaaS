package session

import (
	"defaas/contracts/go/market"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func getSimMarket() (*bind.TransactOpts, *backends.SimulatedBackend, common.Address, *market.MarketSession) {

	auth, blockchain, tokenAddress, _ := getSimFaaSToken()

	address, _, _, _ := market.DeployMarket(auth, blockchain, tokenAddress)
	blockchain.Commit()
	session, _ := NewMarketSeesion(blockchain, address, auth)

	return auth, blockchain, address, session
}

func TestMarketGetFaaSLevel(t *testing.T) {

	auth, blockchain, _, session := getSimMarket()
	_ = auth
	_ = blockchain

	assert := assert.New(t)

	numLevel, _ := session.GetFaaSLevelNumber()
	t.Log("numLevel", numLevel)

	for i := big.NewInt(0); i.Cmp(numLevel) == -1; i.Add(i, big.NewInt(1)) {

		isValid, core, mem, err := session.GetFaaSLevel(i)
		if err != nil {
			log.Fatal(err)
		}

		assert.True(isValid)

		t.Logf("level %v: %v core, %v MB\n", i, core, mem)
	}

	{
		isValid, _, _, err := session.GetFaaSLevel(new(big.Int).Add(numLevel, big.NewInt(1)))
		if err != nil {
			log.Fatal(err)
		}
		assert.False(isValid)
	}
}
