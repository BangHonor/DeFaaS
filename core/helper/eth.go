package helper

import (
	"context"
	"errors"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetETHClient(urls []string) (*ethclient.Client, error) {

	var (
		client *ethclient.Client
		err    error
	)

	if len(urls) == 0 {
		return nil, errors.New("wrong config: empty urls")
	}

	// try connect to some eth node
	for _, url := range urls {

		client, err = ethclient.Dial(url)

		if err == nil {
			return client, nil
		}
	}

	return nil, err
}

func EqualAddress(s, t common.Address) bool {
	// type Address [AddressLength]byte
	// in Go, you can directly compare two arrays of the same length
	return (s == t)
}

func CurrentBlockNumber(client *ethclient.Client) *big.Int {

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return header.Number
}
