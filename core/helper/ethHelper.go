package helper

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

func EqualAddress(s, t common.Address) bool {
	// you can directly compare two arrays of the same length, in Go
	return (s == t)
}
