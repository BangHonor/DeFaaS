package customer

import (
	basic "defaas/client/basic"
	"defaas/core/config"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type CustomerClient struct {
	basic.BasicClient
}

func NewCustomerClient(dfc *config.DeFaaSConfig, key *keystore.Key) (*CustomerClient, error) {

	_basicClient, _ := basic.NewBasicClient(dfc, key)

	client := &CustomerClient{
		BasicClient: *_basicClient,
	}

	return client, nil
}
