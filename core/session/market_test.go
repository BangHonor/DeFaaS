package session

import (
	"log"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarketGetFaaSLevel(t *testing.T) {

	auth, blockchain, _, session := newSimMarket()
	_ = auth
	_ = blockchain

	assert := assert.New(t)

	numLevel, _ := session.NumFaaSLevel()
	t.Log("numLevel", numLevel)

	for i := big.NewInt(0); i.Cmp(numLevel) == -1; i.Add(i, big.NewInt(1)) {

		core, mem, err := session.GetFaaSLevel(i)
		if err != nil {
			log.Fatal(err)
		}

		t.Logf("level %v: %v core, %v MB\n", i, core, mem)
	}

	{
		_, _, err := session.GetFaaSLevel(new(big.Int).Add(numLevel, big.NewInt(1)))
		assert.NotNil(err)
	}
}
