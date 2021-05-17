package helper

import (
	"context"
	"errors"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	NumBlockToWaitRecommended int = 1
)

var (
	ErrWaitPeedingTxTimeout   = errors.New("wait pedding tx timeout")
	ErrWaitMinedBlocksTimeout = errors.New("wait blocks timeout")
)

func waitPeddingTxByPolling(ethClient *ethclient.Client, txHash common.Hash) error {

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

			_, isPending, err := ethClient.TransactionByHash(context.TODO(), txHash)
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

func waitMinedBlocksByPolling(ethClient *ethclient.Client, txHash common.Hash, numBlockToWait int) error {

	// record currnet blockNumber
	curHeader, err := ethClient.HeaderByNumber(context.TODO(), nil)
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
			header, err := ethClient.HeaderByNumber(context.TODO(), nil)
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

func waitPeddingTxBySubscription(ethClient *ethclient.Client, txHash common.Hash) error {

	headers := make(chan *types.Header)
	headerSub, err := ethClient.SubscribeNewHead(context.TODO(), headers)
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

			_, isPending, err := ethClient.TransactionByHash(context.TODO(), txHash)
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

func waitMinedBlocksBySubscription(ethClient *ethclient.Client, txHash common.Hash, numBlockToWait int) error {

	headers := make(chan *types.Header)
	headerSub, err := ethClient.SubscribeNewHead(context.TODO(), headers)
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

func ConfirmTxByPolling(ethClient *ethclient.Client, txHash common.Hash, numBlockToWait int) error {

	log.Printf("[confirmTx] wait pedding tx [%v] ...", txHash)
	if err := waitPeddingTxByPolling(ethClient, txHash); err != nil {
		return err
	}
	log.Printf("[confirmTx] wait pedding tx [%v] done", txHash)

	log.Printf("[confirmTx] wait [%v] mined blocks for tx [%v] ...", numBlockToWait, txHash)
	if err := waitMinedBlocksByPolling(ethClient, txHash, numBlockToWait); err != nil {
		return err
	}
	log.Printf("[confirmTx] wait [%v] mined blocks for tx [%v] done", numBlockToWait, txHash)

	return nil
}

func ConfirmTxBySubscription(ethClient *ethclient.Client, txHash common.Hash, numBlockToWait int) error {

	log.Printf("[confirmTx] wait pedding tx [%v] ...", txHash)
	if err := waitPeddingTxBySubscription(ethClient, txHash); err != nil {
		return err
	}
	log.Printf("[confirmTx] wait pedding tx [%v] done", txHash)

	log.Printf("[confirmTx] wait [%v] mined blocks for tx [%v] ...", numBlockToWait, txHash)
	if err := waitMinedBlocksBySubscription(ethClient, txHash, numBlockToWait); err != nil {
		return err
	}
	log.Printf("[confirmTx] wait [%v] mined blocks for tx [%v] done", numBlockToWait, txHash)

	return nil
}
