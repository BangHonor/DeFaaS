package basic

import (
	"context"
	"defaas/core/config"
	"defaas/core/helper"
	"defaas/core/suite"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

func init() {
	log.SetFlags(log.Llongfile)
}

type BasicClient struct {
	suite.Suite
	Key          *keystore.Key
	DeFaaSConfig *config.DeFaaSConfig
	ETHClient    *ethclient.Client
}

func NewBasicClient(dfc *config.DeFaaSConfig, key *keystore.Key) (*BasicClient, error) {

	var (
		err error
	)

	log.Println("[basic] ------------------- new basic client .... --------------------------")

	client := &BasicClient{}
	client.Key = key
	client.DeFaaSConfig = dfc

	// connect to eth network
	client.ETHClient, err = helper.GetETHClient(dfc.WsURLs)
	if err != nil {
		return nil, err
	}
	log.Println("[basic] connect to blockchain network")

	// get chainID
	chainID, err := client.ETHClient.NetworkID(context.TODO())
	if err != nil {
		return nil, err
	}
	log.Println("[basic] chainID", chainID)

	// build a auth
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		return nil, err
	}

	_suite, err := suite.NewSuite(
		client.ETHClient,
		auth,
		dfc.FaaSTokenContractAddress,
		dfc.MarketContractAddress,
		dfc.WitnessPoolContractAddress)
	if err != nil {
		return nil, err
	}
	client.Suite = *_suite
	log.Println("[basic] set up a session of smart contarct")

	return client, nil
}

func NewBasicClientWithFile(configFilePath, keyStoreFilePath, password string) (*BasicClient, error) {

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
