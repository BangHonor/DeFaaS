package customer

import (
	basic "defaas/client/basic"
	"defaas/core/config"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type CustomerClient struct {
	basic.BasicClient
}

func NewCustomerClientWithFile(configFilePath, keyStoreFilePath string, password string) (*CustomerClient, error) {

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

	return NewCustomerClient(dfc, key)
}

func NewCustomerClient(dfc *config.DeFaaSConfig, key *keystore.Key) (*CustomerClient, error) {

	_basicClient, _ := basic.NewBasicClient(dfc, key)

	client := &CustomerClient{
		BasicClient: *_basicClient,
	}

	return client, nil
}
