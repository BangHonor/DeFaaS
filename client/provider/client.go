package provider

import (
	"context"
	"defaas/client/basic"
	"defaas/core/config"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gogf/gf/container/gmap"
)

type ProviderClient struct {
	basic.BasicClient

	accessKeyPool *gmap.AnyAnyMap

	sinkBiddenID chan *big.Int

	quit context.CancelFunc
}

func NewProviderClient(dfc *config.DeFaaSConfig, key *keystore.Key) (*ProviderClient, error) {

	client := &ProviderClient{}

	_basicClient, err := basic.NewBasicClient(dfc, key)
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

	log.Println("[provider] start bidding ...")

	if err := client.Bidding(ctx); err != nil {
		return err
	}

	log.Panicln("[provider] start bidding done")

	log.Println("[provider] start deploy server ...")

	// go client.StartDeployServer()

	log.Panicln("[provider] start deploy serve done")

	return nil
}

func (client *ProviderClient) Close() error {

	client.quit()

	return nil
}
