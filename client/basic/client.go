package basic

import (
	"context"
	"defaas/core/config"
	"defaas/core/helper"
	"defaas/core/session"
	"io/ioutil"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BasicClient struct {
	rawClinet   *ethclient.Client
	faastoken   *session.FaaSTokenSession
	market      *session.MarketSeesion
	witnesspool *session.WitnessPoolSession

	key *keystore.Key
}

func NewBasicClientWithFile(configFilePath, keyStoreFilePath string, password string) (*BasicClient, error) {

	// parse defaas config
	dfc, err := config.ParseConfigFile(configFilePath)
	if err != nil {
		return nil, err
	}

	// decrypt keystore
	keyjson, err := ioutil.ReadFile(keyStoreFilePath)
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(keyjson, password)
	if err != nil {
		return nil, err
	}

	return NewBasicClient(dfc, key)
}

func NewBasicClient(dfc *config.DeFaaSConfig, key *keystore.Key) (*BasicClient, error) {

	var (
		err error
	)

	client := &BasicClient{}
	client.key = key

	// connect to eth network
	client.rawClinet, err = helper.GetETHClient(dfc.WsURLs)
	if err != nil {
		return nil, err
	}

	// get chainID
	chainID, err := client.rawClinet.NetworkID(context.TODO())
	if err != nil {
		return nil, err
	}

	// build a auth
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		return nil, err
	}

	// build session
	client.faastoken, err = session.NewFaaSTokenSeesion(client.rawClinet, dfc.FaaSTokenContractAddress, auth)
	if err != nil {
		return nil, err
	}

	client.market, err = session.NewMarketSeesion(client.rawClinet, dfc.MarketContractAddress, auth)
	if err != nil {
		return nil, err
	}

	client.witnesspool, err = session.NewWitnessPoolSession(client.rawClinet, dfc.WitnessPoolContractAddress, auth)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (client *BasicClient) BalanceOf(address common.Address) (*big.Int, error) {
	return big.NewInt(0), nil
}
