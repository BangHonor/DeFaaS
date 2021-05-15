package basic

import (
	"context"
	"errors"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	NumBlockToWaitRecommended int = 1
)

var (
	ErrWaitPeedingTxTimeout   = errors.New("wait pedding tx timeout")
	ErrWaitMinedBlocksTimeout = errors.New("wait blocks timeout")
)

func (client *BasicClient) waitPeddingTxByPolling(txHash common.Hash) error {

	checkPeddingTimeout := time.NewTimer(1 * time.Minute)
	checkPeddingTicker := time.NewTicker(1 * time.Second)
	defer checkPeddingTicker.Stop()

	// wait for pedding
CheckPeddingLoop:

	for {
		select {

		case <-checkPeddingTimeout.C:
			return ErrWaitPeedingTxTimeout

		case <-checkPeddingTicker.C:

			_, isPending, err := client.ETHClinet.TransactionByHash(context.TODO(), txHash)
			if err != nil {
				return err
			}

			if !isPending {
				break CheckPeddingLoop
			}
		}
	}

	return nil
}

func (client *BasicClient) waitMinedBlocksByPolling(txHash common.Hash, numBlockToWait int) error {

	// record currnet blockNumber
	curHeader, err := client.ETHClinet.HeaderByNumber(context.TODO(), nil)
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
			return ErrWaitMinedBlocksTimeout
		case <-waitBlockTicker.C:
			header, err := client.ETHClinet.HeaderByNumber(context.TODO(), nil)
			if err != nil {
				return err
			}
			if big.NewInt(int64(numBlockToWait)).Cmp(header.Number.Sub(header.Number, curHeader.Number)) <= 0 {
				break WaitBlockLoop
			}
		}
	}

	return nil
}

func (client *BasicClient) waitPeddingTxBySubscription(txHash common.Hash) error {

	headers := make(chan *types.Header)
	headerSub, err := client.ETHClinet.SubscribeNewHead(context.TODO(), headers)
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

			_, isPending, err := client.ETHClinet.TransactionByHash(context.TODO(), txHash)
			if err != nil {
				return err
			}

			if !isPending {
				break CheckPeddingSubLoop
			}

		}
	}

	return nil
}

func (client *BasicClient) waitMinedBlocksBySubscription(txHash common.Hash, numBlockToWait int) error {

	headers := make(chan *types.Header)
	headerSub, err := client.ETHClinet.SubscribeNewHead(context.TODO(), headers)
	if err != nil {
		return err
	}

	// wait for blocks
	for i := 0; i < numBlockToWait; i++ {
		select {
		case err := <-headerSub.Err():
			return err
		case <-headers:
			// do nothig
		}
	}

	return nil
}

func (client *BasicClient) ComfirmTxByPolling(txHash common.Hash, numBlockToWait int) error {

	log.Printf("[basic] wait pedding tx [%v] ...", txHash)
	if err := client.waitPeddingTxByPolling(txHash); err != nil {
		return err
	}
	log.Printf("[basic] wait pedding tx [%v] done", txHash)

	log.Printf("[basic] wait [%v] mined blocks for tx [%v] ...", numBlockToWait, txHash)
	if err := client.waitMinedBlocksByPolling(txHash, numBlockToWait); err != nil {
		return err
	}
	log.Printf("[basic] wait [%v] mined blocks for tx [%v] done", numBlockToWait, txHash)

	return nil
}

func (client *BasicClient) ComfirmTxBySubscription(txHash common.Hash, numBlockToWait int) error {

	log.Printf("[basic] wait pedding tx [%v] ...", txHash)
	if err := client.waitPeddingTxBySubscription(txHash); err != nil {
		return err
	}
	log.Printf("[basic] wait pedding tx [%v] done", txHash)

	log.Printf("[basic] wait [%v] mined blocks for tx [%v] ...", numBlockToWait, txHash)
	if err := client.waitMinedBlocksBySubscription(txHash, numBlockToWait); err != nil {
		return err
	}
	log.Printf("[basic] wait [%v] mined blocks for tx [%v] done", numBlockToWait, txHash)

	return nil
}
