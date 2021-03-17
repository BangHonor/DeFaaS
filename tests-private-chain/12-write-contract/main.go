package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"defaas/contracts/go/faastoken"
)

func main() {

	// 构造一个连接
	client, err := ethclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatal(err)
	}

	// 获取 chainID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chainID", chainID)

	// 加载一个账户的私钥
	privateKeyFile := "./tmp/dev/data/keystore/UTC--2021-03-15T08-50-30.516148546Z--f38f26975aec981583ae8e4029f640c1b0d7f91a"
	password := ""
	keyjson, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	// 构造一个授权事务
	auth, err := bind.NewTransactorWithChainID(bytes.NewBuffer(keyjson), password, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	contractAddress := common.HexToAddress("0x26e0fa32260d3fc55fd5f3292bc7de35f5d8592f")
	instance, err := faastoken.NewFaaSToken(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("bytecode len", len(hex.EncodeToString(bytecode)))

	session := faastoken.FaaSTokenSession{
		Contract: instance,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:    auth.From,
			Signer:  auth.Signer,
			Context: context.TODO(),
		},
	}

	name, err := session.Name()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("token name", name)

	queryAddress := common.HexToAddress("f38f26975aec981583ae8e4029f640c1b0d7f91a")

	balanceOf, err := session.BalanceOf(queryAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balanceOf", queryAddress.Hex(), balanceOf)

	tx, err := session.Mint(big.NewInt(100))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx sent:", tx.Hash().Hex())

	balanceOf, err = session.BalanceOf(queryAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balanceOf", queryAddress.Hex(), balanceOf)

	tx, err = session.MintTo(common.BigToAddress(big.NewInt(0)), big.NewInt(100))
	if err == nil {
		log.Fatal(errors.New("expected error"))
	}
	fmt.Println(err)
	fmt.Println(tx)
}
