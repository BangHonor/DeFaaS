package basic

import (
	"context"
	"defaas/core/config"
	"defaas/core/helper"
	"defaas/core/session"
	"errors"
	"io/ioutil"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	NumBlockToWaitRecommended int = 1
)

type BasicClient struct {
	Key          *keystore.Key
	DeFaaSConfig *config.DeFaaSConfig

	RawClinet   *ethclient.Client
	FaaSToken   *session.FaaSTokenSession
	Market      *session.MarketSeesion
	WitnessPool *session.WitnessPoolSession
}

func NewBasicClientWithFile(configFilePath, keyStoreFilePath string, password string) (*BasicClient, error) {

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

	return NewBasicClient(dfc, key)
}

func NewBasicClient(dfc *config.DeFaaSConfig, key *keystore.Key) (*BasicClient, error) {

	log.Println("[basic] ------------------- new basic client .... --------------------------")

	var (
		err error
	)

	client := &BasicClient{}
	client.Key = key
	client.DeFaaSConfig = dfc

	// connect to eth network
	client.RawClinet, err = helper.GetETHClient(dfc.WsURLs)
	if err != nil {
		return nil, err
	}
	log.Println("[basic] connect to blockchain network")

	// get chainID
	chainID, err := client.RawClinet.NetworkID(context.TODO())
	if err != nil {
		return nil, err
	}
	log.Println("[basic] chainID", chainID)

	// build a auth
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		return nil, err
	}

	// build session
	client.FaaSToken, err = session.NewFaaSTokenSeesion(client.RawClinet, dfc.FaaSTokenContractAddress, auth)
	if err != nil {
		return nil, err
	}
	log.Println("[basic] set up a session of [FaaSToekn] contarct")

	client.Market, err = session.NewMarketSeesion(client.RawClinet, dfc.MarketContractAddress, auth)
	if err != nil {
		return nil, err
	}
	log.Println("[basic] set up a session of [Market] contarct")

	client.WitnessPool, err = session.NewWitnessPoolSession(client.RawClinet, dfc.WitnessPoolContractAddress, auth)
	if err != nil {
		return nil, err
	}
	log.Println("[basic] set up a session of [WitnessPool] contarct")

	log.Println("[basic] ------------------- new basic client done --------------------------")

	return client, nil
}

func (client *BasicClient) ComfirmTxByPolling(txHash common.Hash, numBlockToWait int) error {

	log.Printf("[basic] comfirm tx [%v] for pedding ...", txHash)

	checkPeddingTimeout := time.NewTimer(1 * time.Minute)
	checkPeddingTicker := time.NewTicker(1 * time.Second)
	defer checkPeddingTimeout.Stop()
	defer checkPeddingTicker.Stop()

	// wait for pedding
CheckPeddingLoop:

	for {
		select {

		case <-checkPeddingTimeout.C:
			return errors.New("chekc pedding time out")

		case <-checkPeddingTicker.C:

			_, isPending, err := client.RawClinet.TransactionByHash(context.TODO(), txHash)
			if err != nil {
				return err
			}

			if !isPending {
				break CheckPeddingLoop
			}
		}
	}

	log.Printf("[basic] comfirm tx [%v] for pedding done", txHash)

	log.Printf("[basic] comfirm tx [%v] for waiting [%v] blocks ...", txHash, numBlockToWait)

	// record currnet blockNumber
	curHeader, err := client.RawClinet.HeaderByNumber(context.TODO(), nil)
	if err != nil {
		return err
	}

	waitBlockTimeout := time.NewTimer(1 * time.Minute)
	waitBlockTicker := time.NewTicker(1 * time.Second)
	defer waitBlockTimeout.Stop()
	defer waitBlockTicker.Stop()

	// wait fot blocks
WaitBlockLoop:

	for {
		select {
		case <-waitBlockTimeout.C:
			return errors.New("wait block time out")
		case <-waitBlockTicker.C:
			header, err := client.RawClinet.HeaderByNumber(context.TODO(), nil)
			if err != nil {
				return err
			}
			if big.NewInt(int64(numBlockToWait)).Cmp(header.Number.Sub(header.Number, curHeader.Number)) <= 0 {
				break WaitBlockLoop
			}
		}
	}

	log.Printf("[basic] comfirm tx [%v] for waiting [%v] blocks done", txHash, numBlockToWait)

	return nil
}

func (client *BasicClient) ComfirmTxBySubscription(txHash common.Hash, numBlockToWait int) error {

	log.Printf("[basic] comfirm tx [%v] for pedding ...", txHash)

	headers := make(chan *types.Header)
	headerSub, err := client.RawClinet.SubscribeNewHead(context.TODO(), headers)
	if err != nil {
		return err
	}

	checkPeddingTimeout := time.NewTimer(1 * time.Minute)

	// wait for pedding
CheckPeddingSubLoop:
	for {
		select {

		case <-checkPeddingTimeout.C:
			return errors.New("chekc pedding time out")

		case err := <-headerSub.Err():
			return err

		case <-headers:

			_, isPending, err := client.RawClinet.TransactionByHash(context.TODO(), txHash)
			if err != nil {
				return err
			}

			if !isPending {
				break CheckPeddingSubLoop
			}

		}
	}

	log.Printf("[basic] comfirm tx [%v] for pedding done", txHash)

	log.Printf("[basic] comfirm tx [%v] for waiting [%v] blocks ...", txHash, numBlockToWait)

	// wait for blocks
	for i := 0; i < numBlockToWait; i++ {
		<-headers
	}

	log.Printf("[basic] comfirm tx [%v] for waiting [%v] blocks done", txHash, numBlockToWait)

	return nil
}
