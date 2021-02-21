package main

// 文档
// https://github.com/FISCO-BCOS/go-sdk/blob/master/doc/README.md#%E7%8E%AF%E5%A2%83%E9%85%8D%E7%BD%AE#%E5%A4%96%E9%83%A8%E8%B4%A6%E6%88%B7

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func generateExample() {
	// 生成随机私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// 转换私钥为字节序列
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("private key: ", hexutil.Encode(privateKeyBytes)[2:]) // privateKey in hex without "0x"

	// 公钥是从私钥派生的，根据加密私钥返回公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 转换公钥为字节序列
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publick key: ", hexutil.Encode(publicKeyBytes)[4:]) // publicKey in hex without "0x"

	// 转换公钥为地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("address: ", address.Hex()) // account address

}

func configExample() {
	// load config
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]

	fmt.Println("config address: ", getAddressFromConfig(config).Hex())
}

func getAddressFromConfig(config *conf.Config) common.Address {
	privateKey, err := crypto.ToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal("parse error")
	}
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}

func main() {
	generateExample()
	configExample()
}
