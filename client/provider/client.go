package provider

import (
	"context"
	"defaas/client/basic"
	"math/big"

	"github.com/gogf/gf/container/gmap"
)

type ProviderClient struct {
	basic.BasicClient

	accessKeyPool *gmap.AnyAnyMap

	sinkBiddenID chan *big.Int

	quit context.CancelFunc
}

func NewProviderClient() (*ProviderClient, error) {

	client := &ProviderClient{}

	_basicClient, err := basic.NewBasicClient(nil, nil)
	if err != nil {
		return nil, err
	}
	client.BasicClient = *_basicClient

	client.accessKeyPool = gmap.NewHashMap(true) // `trur` means concurrent-safety

	client.sinkBiddenID = make(chan *big.Int)

	return client, nil
}

func (client *ProviderClient) Start() error {

	ctx, cancel := context.WithCancel(context.Background())
	client.quit = cancel

	if err := client.Bidding(ctx); err != nil {
		return err
	}

	// go client.StartDeployServer()

	return nil
}

func (client *ProviderClient) Close() error {

	client.quit()

	return nil
}
