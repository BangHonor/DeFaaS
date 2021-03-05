package main

// 文档
// https://github.com/FISCO-BCOS/go-sdk/blob/master/doc/README.md#%E7%8E%AF%E5%A2%83%E9%85%8D%E7%BD%AE#%E5%A4%96%E9%83%A8%E8%B4%A6%E6%88%B7

import (
	"fmt"
	"log"

	"defaas/dev-cmd/utils"

	"github.com/FISCO-BCOS/go-sdk/conf"
)

// func generateExample() {
// 	// 生成随机私钥
// 	privateKey, err := crypto.GenerateKey()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 转换私钥为字节序列
// 	privateKeyBytes := crypto.FromECDSA(privateKey)
// 	fmt.Println("private key: ", hexutil.Encode(privateKeyBytes)[2:]) // privateKey in hex without "0x"

// 	// 公钥是从私钥派生的，根据加密私钥返回公钥
// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
// 	}

// 	// 转换公钥为字节序列
// 	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
// 	fmt.Println("publick key: ", hexutil.Encode(publicKeyBytes)[4:]) // publicKey in hex without "0x"

// 	// 转换公钥为地址
// 	address := crypto.PubkeyToAddress(*publicKeyECDSA)
// 	fmt.Println("generate address: ", address.Hex()) // account address
// }

func generateExample() {
	// generate
	address, _ := utils.GetAddressFromGenerating()
	fmt.Println("generate address: ", address.Hex()) // account address
}

func configExample() {
	// load config
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]

	address, _ := utils.GetAddressFromConfig(config)
	fmt.Println("generate address: ", address.Hex()) // account address
}

func pemExample() {
	pem := "./accounts/0x9db2cf0827d258fb96a81f57839747c27a31a4d2.pem"
	address, _ := utils.GetAddressFromPem(pem)
	fmt.Println("pem      address: ", address.Hex())
}

func main() {
	generateExample()
	configExample()
	pemExample()
}
