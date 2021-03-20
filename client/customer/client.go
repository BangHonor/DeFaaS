package customer

import (
	basic "defaas/client/basic"
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
