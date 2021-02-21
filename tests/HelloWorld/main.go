package main

import (
	"fmt"
	"log"

	// "defaas/helloworld" // import helloworld
	"defaas/contracts/go/helloworld" // import helloworld

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
)

func main() {

	// load config
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]

	// client connect...
	client, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}

	// deploy the contract
	address, tx, _instance, err := helloworld.DeployHelloWorld(client.GetTransactOpts(), client) // deploy contract
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract address: ", address.Hex()) // the address should be saved
	fmt.Println("transaction hash: ", tx.Hash().Hex())
	_ = _instance

	// load the contract
	contractAddress := common.HexToAddress(address.Hex())
	instance, err := helloworld.NewHelloWorld(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	helloworldSession := &helloworld.HelloWorldSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	value, err := helloworldSession.Get() // call Get API
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value before set:", value)

	value = "Hello, FISCO BCOS"
	tx, receipt, err := helloworldSession.Set(value) // call set API
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("set value:", value)
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

	value, err = helloworldSession.Get()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value after set:", value)
}
