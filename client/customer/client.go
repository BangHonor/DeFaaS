package customer

import (
	basic "defaas/client/basic"

	"github.com/ethereum/go-ethereum/common"
)

type CustomerClient struct {
	basic.BasicClient
}

func NewCustomerClient() (*CustomerClient, error) {

	_basicClient, _ := basic.NewBasicClient(nil, nil)

	client := &CustomerClient{
		BasicClient: *_basicClient,
	}

	return client, nil
}

func (client *CustomerClient) Deploy() error {
	client.FaaSToken.BalanceOf(common.HexToAddress("0x01"))
	return nil
}
