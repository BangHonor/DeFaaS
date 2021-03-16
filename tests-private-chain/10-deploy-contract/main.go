package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"defaas/contracts/go/faastoken"
)

func main() {

	// 构造一个连接
	client, err := ethclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatal(err)
	}

	// 加载一个账户的私钥
	privateKeyFile := "./tmp/dev/data/keystore/UTC--2021-03-15T08-50-30.516148546Z--f38f26975aec981583ae8e4029f640c1b0d7f91a"
	password := ""
	keyjson, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	// 获取 chainID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(chainID)

	// 构造一个授权事务
	auth, err := bind.NewTransactorWithChainID(bytes.NewBuffer(keyjson), password, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// 部署
	address, tx, instance, err := faastoken.DeployFaaSToken(auth, client)
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	totolSupply, err := instance.TotalSupply(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("totolSupply:", totolSupply)
}
