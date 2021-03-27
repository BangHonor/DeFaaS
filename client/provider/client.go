package provider

import (
	"context"
	"defaas/client/basic"
	defaasconfig "defaas/core/config"
	"defaas/core/data"
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

	client.deployServer = ghttp.GetServer("deploy-server")
	client.deployServer.SetAddr(data.ServerAddrTrim(client.providerConfig.ServerAddr))
	client.deployServer.BindHandler(client.providerConfig.ServerEntry, client.DeployServerRequestHandler)

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
	bidder2fulfiller := make(chan *data.DeploymentItem, 1)

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
