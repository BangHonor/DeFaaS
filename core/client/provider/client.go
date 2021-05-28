package provider

import (
	"context"
	"defaas/core/client/basic"
	defaasconfig "defaas/core/config"
	model "defaas/core/model"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gogf/gf/container/gmap"
)

type ProviderClient struct {
	basic.BasicClient

	itemPool *gmap.AnyAnyMap // map[*big.Int]*model.DeploymentItem, map: ID => Item

	quit context.CancelFunc
}

func NewProviderClientWithFile(defaasConfigFilePath, keyStoreFilePath, password string) (*ProviderClient, error) {

	// parse defaas config
	dfc, err := defaasconfig.ParseConfigFile(defaasConfigFilePath)
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

	return NewProviderClient(dfc, key)
}

func NewProviderClient(dfc *defaasconfig.DeFaaSConfig, key *keystore.Key) (*ProviderClient, error) {

	client := &ProviderClient{}

	_basicClient, err := basic.NewBasicClient(dfc, key)
	if err != nil {
		return nil, err
	}
	client.BasicClient = *_basicClient

	client.itemPool = gmap.NewHashMap(true) // `true` means concurrent-safety

	return client, nil
}

func (client *ProviderClient) Start() error {

	var (
		err error = nil
	)

	ctx, cancel := context.WithCancel(context.Background())
	client.quit = cancel

	defer func() {
		if err != nil {
			log.Println("[provider] an error occurred while starting the client")
			if err := client.Close(); err != nil {
				log.Fatal("[provider] fail to quit safely when an error starting the client")
			}
			log.Println("[provider] quit safely when an error starting the client")
		}
	}()

	watcher2bidder := make(chan *model.DeploymentItem, 1)
	bidder2fulfiller := make(chan *model.DeploymentItem, 1)

	watcherErr := make(chan error)
	bidderErr := make(chan error)
	fulfillerErr := make(chan error)

	log.Println("[provider] start [watcher] ...")
	if err = client.Watcher(ctx, watcher2bidder, watcherErr); err != nil {
		log.Println("[provider] failed to start [watcher]")
		return err
	}
	log.Println("[provider] start [watcher] done")

	log.Println("[provider] start [bidder] ...")
	if err = client.Bidder(ctx, watcher2bidder, bidder2fulfiller, bidderErr); err != nil {
		log.Println("[provider] failed to start [bidder]")
		return err
	}
	log.Println("[provider] start [bidder] done")

	log.Println("[provider] start [fulfiller] ...")
	if err = client.Fulfiller(ctx, bidder2fulfiller, fulfillerErr); err != nil {
		log.Println("[provider] failed to start [fulfiller]")
		return err
	}
	log.Println("[provider] start [fulfiller] done")

	// watch errors in the worker goroutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-watcherErr:
				_ = err
			case err := <-bidderErr:
				_ = err
			case err := <-fulfillerErr:
				_ = err
			}
		}
	}()

	return nil
}

func (client *ProviderClient) Close() error {

	log.Println("[provider] closing ...")

	client.quit()

	return nil
}
