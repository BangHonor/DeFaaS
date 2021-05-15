package customer

import (
	"crypto/sha256"
	basic "defaas/core/client/basic"
	"defaas/core/config"
	"defaas/core/data"
	"io/ioutil"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gogf/gf/container/gmap"
)

type CustomerClient struct {
	basic.BasicClient

	itemPool                   *gmap.AnyAnyMap // map[big.Int]data.DeploymentItem, map: ID => Item
	nonce                      int64
	privateKey                 data.PrivateKey
	biddingDurationRecommended *big.Int
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

	client := &CustomerClient{}

	_basicClient, err := basic.NewBasicClient(dfc, key)
	if err != nil {
		return nil, err
	}
	client.BasicClient = *_basicClient

	client.itemPool = gmap.NewHashMap(true) // `true` means concurrent-safety

	client.nonce = 0

	privateKey, err := data.GeneratePrivateKey()
	if err != nil {
		return nil, err
	}
	client.privateKey = privateKey

	client.biddingDurationRecommended = big.NewInt(int64(time.Minute.Seconds())) // 1 minute

	return client, nil
}

func (client *CustomerClient) NewDeploy(faasLevelID, faasDuration, highestUnitPrice *big.Int, adapterData []byte) error {

	order, err := client.newOrder(faasLevelID, highestUnitPrice, faasDuration)
	if err != nil {
		return err
	}

	err = client.DeployOrder(order, adapterData)
	if err != nil {
		return err
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (client *CustomerClient) newOrder(faasLevelID, highestUnitPrice, faasDuration *big.Int) (*data.DeploymentOrder, error) {

	fulfillSecretKey := data.GenerateFulfillSecretKey()

	encodedPublicKey, err := data.EncodePublicKey(client.privateKey.Public())
	if err != nil {
		return nil, err
	}

	order := &data.DeploymentOrder{
		Customer:         client.Key.Address,
		Nonce:            big.NewInt(atomic.AddInt64(&client.nonce, 1)),
		FaaSLevelID:      faasLevelID,
		HighestUnitPrice: highestUnitPrice,
		FaaSDuration:     faasDuration,
		BiddingDuration:  client.biddingDurationRecommended,
		PublicKey:        string(encodedPublicKey),
		FulfillSecretKey: fulfillSecretKey,
		FulfillKey:       sha256.Sum256(fulfillSecretKey[:]),
	}

	return order, nil
}
