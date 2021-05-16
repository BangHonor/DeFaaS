package customer

import (
	"defaas/core/client/basic"
	"defaas/core/config"
	"io/ioutil"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gogf/gf/container/gmap"
)

var (
	BiddingDurationRecommended = big.NewInt(int64(time.Minute.Seconds())) // 1 minute
)

type CustomerClient struct {
	basic.BasicClient
	itemPool *gmap.AnyAnyMap // map[*big.Int]*model.DeploymentItem, map: ID => Item
	nonce    int64
}

func NewCustomerClientWithFile(configFilePath, keyStoreFilePath, password string) (*CustomerClient, error) {

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

	client := &CustomerClient{}

	_basicClient, err := basic.NewBasicClient(dfc, key)
	if err != nil {
		return nil, err
	}
	client.BasicClient = *_basicClient

	client.itemPool = gmap.NewHashMap(true) // `true` means concurrent-safety

	client.nonce = 0

	return client, nil
}
