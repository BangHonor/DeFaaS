package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func init() {
	// log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetFlags(log.Lmicroseconds | log.Llongfile)
}

func f1() {
	address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	fmt.Println(address.Hex())        // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
	fmt.Println(address.Hash().Hex()) // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
	fmt.Println(address.Bytes())      // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]

}

func f2() {
	// client, err := ethclient.Dial("https://goerli.infura.io/v3/1d69d7036ac046af9e31ff2a789d74c0")

	client, err := ethclient.Dial("http://127.0.0.1:8545")

	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
}

func f3() {

	// 默认是 ipc
	// client, err := ethclient.Dial("/home/kitchen/geth-learn/private_chain/data-0/geth.ipc")

	client, err := ethclient.Dial("http://127.0.0.1:8545")

	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x8df69a21a6bde726cf249b4f792784215a0436e7")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	fromAddress := "0xf38f26975aec981583ae8e4029f640c1b0d7f91a"
	account = common.HexToAddress(fromAddress)
	balance, err = client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance, "of", fromAddress)

	toAddress := "0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d"
	account = common.HexToAddress(toAddress)
	balance, err = client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance, "of", toAddress)

}

func f4() {

	fmt.Println("ws")
	client, err := ethclient.Dial("ws://127.0.0.1:8546")

	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
}

func main() {
	f1()
	f2()
	f3()
	f4()
}
