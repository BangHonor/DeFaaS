package provider

import (
	"context"
	"defaas/client/basic"
	defaasconfig "defaas/core/config"
	"defaas/core/data"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
)

type ProviderClient struct {
	basic.BasicClient

	providerConfig *ProviderConfig

	itemPool             *gmap.AnyAnyMap // map[big.Int]data.DeploymentItem, map: ID => Item
	deployedNotifierPool *gmap.AnyAnyMap // map[big.Int](chan *data.DeploymentItem), map: ID => notifier

	quit context.CancelFunc

	deployServer *ghttp.Server
}

func NewProviderClientWithFile(defaasConfigFilePath, providerConfigFilePath, keyStoreFilePath string, password string) (*ProviderClient, error) {

	// parse defaas config
	dfc, err := defaasconfig.ParseConfigFile(defaasConfigFilePath)
	if err != nil {
		return nil, err
	}

	// parse provider config
	pc, err := ParseConfigFile(providerConfigFilePath)
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

	return NewProviderClient(dfc, pc, key)
}

func NewProviderClient(dfc *defaasconfig.DeFaaSConfig, pc *ProviderConfig, key *keystore.Key) (*ProviderClient, error) {

	client := &ProviderClient{}

	_basicClient, err := basic.NewBasicClient(dfc, key)
	if err != nil {
		return nil, err
	}
	client.BasicClient = *_basicClient

	client.providerConfig = pc

	client.itemPool = gmap.NewHashMap(true) // `true` means concurrent-safety

	client.deployServer = ghttp.GetServer("deploy-server-" + fmt.Sprint(0))
	client.deployServer.SetAddr(client.providerConfig.ServerAddr)                                         // https://pkg.go.dev/github.com/gogf/gf/net/ghttp#Server.SetAddr
	client.deployServer.BindHandler(client.providerConfig.ServerEntry, client.DeployServerRequestHandler) // register handler

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

	watcher2bidder := make(chan *data.DeploymentItem, 1)
	bidder2publisher := make(chan *data.DeploymentItem, 1)
	publisher2fulfiller := make(chan *data.DeploymentItem, 1)
	fulfiller2finisher := make(chan *data.DeploymentItem, 1)

	watcherErr := make(chan error)
	bidderErr := make(chan error)
	publisherErr := make(chan error)
	fulfillerErr := make(chan error)
	finisherErr := make(chan error)

	log.Println("[provider] start [watcher] ...")
	if err = client.Watcher(ctx, watcher2bidder, watcherErr); err != nil {
		log.Println("[provider] start [watcher] failed")
		return err
	}
	log.Println("[provider] start [watcher] done")

	log.Println("[provider] start [bidder] ...")
	if err = client.Bidder(ctx, watcher2bidder, bidder2publisher, bidderErr); err != nil {
		log.Println("[provider] start [bidder] failed")
		return err
	}
	log.Println("[provider] start [bidder] done")

	log.Println("[provider] start [publisher] ...")
	if err = client.Publisher(ctx, bidder2publisher, publisher2fulfiller, publisherErr); err != nil {
		log.Println("[provider] start [publisher] failed")
		return err
	}
	log.Println("[provider] start [publisher] done")

	log.Println("[provider] start [fulfiller] ...")
	if err = client.Fulfiller(ctx, publisher2fulfiller, fulfiller2finisher, fulfillerErr); err != nil {
		log.Println("[provider] start [fulfiller] failed")
		return err
	}
	log.Println("[provider] start [fulfiller] done")

	log.Println("[provider] start [finisher] ...")
	if err = client.Finisher(ctx, fulfiller2finisher, finisherErr); err != nil {
		log.Println("[provider] start [finisher] failed")
		return err
	}
	log.Println("[provider] start [finisher] done")

	go func() {
		for {
			select {

			case <-ctx.Done():
				return

			case err := <-watcherErr:
				log.Println("[provider/watcher]", err)

			case err := <-bidderErr:
				log.Println("[provider/bidder]", err)

			case err := <-publisherErr:
				log.Println("[provider/publisher]", err)

			case err := <-fulfillerErr:
				log.Println("[provider/fulfiller]", err)

			case err := <-finisherErr:
				log.Println("[provider/finisher]", err)

			}
		}
	}()

	log.Println("[provider] start deploy server ...")
	client.deployServer.Run()
	log.Printf("[provider] run deploy server [%s]\n", GetDeployPathFromProviderConfig(client.providerConfig))
	log.Panicln("[provider] start deploy serve done")

	return nil
}

func (client *ProviderClient) Close() error {

	log.Println("[provider] closing ...")

	client.quit()

	err := client.deployServer.Shutdown()
	if err != nil {
		log.Printf("[provider] failed to shutdown deploy server [%s]\n", GetDeployPathFromProviderConfig(client.providerConfig))
		return err
	}
	log.Printf("[provider] shutdown deploy server [%s]\n", GetDeployPathFromProviderConfig(client.providerConfig))

	return nil
}
