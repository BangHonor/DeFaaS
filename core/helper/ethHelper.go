package helper

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 返回一个授权事务
func GetAuthFromKeyStore(keyStoreFilePath, password string, client *ethclient.Client) (*bind.TransactOpts, error) {

	// 获取 chainID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	// 加载私钥
	keyjson, err := ioutil.ReadFile(keyStoreFilePath)
	if err != nil {
		return nil, err
	}

	// 构造一个授权事务
	auth, err := bind.NewTransactorWithChainID(bytes.NewBuffer(keyjson), password, chainID)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func GetETHClient(urls []string) (*ethclient.Client, error) {

	var (
		client *ethclient.Client = nil
		err    error             = nil
	)

	if len(urls) == 0 {
		return nil, errors.New("wrong config: empty urls")
	}

	for _, url := range urls {

		client, err = ethclient.Dial(url)
		if err == nil {
			return client, nil
		}
	}

	return nil, err
}

// 新建一个模拟区块链
func NewSim() (*bind.TransactOpts, *backends.SimulatedBackend) {

	// Generate a new random account and a funded simulator
	key, _ := crypto.GenerateKey()
	auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 1<<32)

	return auth, blockchain
}

func EqualAddress(s, t common.Address) bool {
	// you can directly compare two arrays of the same length, in Go
	return (s == t)
}
