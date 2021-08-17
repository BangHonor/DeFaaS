package bigintmap

import (
	"log"
	"math/big"
	"testing"

	"github.com/gogf/gf/container/gmap"
	"github.com/stretchr/testify/assert"
)

func TestBigIntMap(t *testing.T) {

	assert := assert.New(t)

	bi := big.NewInt(3)

	testKVs := []struct {
		k *big.Int
		v string
	}{
		{
			big.NewInt(0),
			"0",
		},
		{
			big.NewInt(1),
			"1",
		},
		{
			big.NewInt(233),
			"233",
		}, {
			bi,
			"3-1",
		},
		{
			bi,
			"3-2",
		},
		{
			bi,
			"3-3",
		},
	}

	m := gmap.NewHashMap(true)

	for _, kv := range testKVs {
		m.Set(kv.k, kv.v)
	}

	m.Iterator(func(key, value interface{}) bool {
		log.Println(key, value)
		return true
	})

	assert.Nil(m.Get(big.NewInt(555)))

}
