package utils

// 产生地址

import (
	"fmt"
	"log"

	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// GetAddressFromConfig ...
func GetAddressFromConfig(config *conf.Config) (address common.Address, privateKeyBytes []byte) {
	privateKey, err := crypto.ToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal("parse error")
	}
	return crypto.PubkeyToAddress(privateKey.PublicKey), config.PrivateKey
}

// GetAddressFromPem ...
func GetAddressFromPem(pemFilePath string) (address common.Address, privateKeyBytes []byte) {

	privateKeyBytes, curve, err := conf.LoadECPrivateKeyFromPEM(pemFilePath)

	if err != nil {
		log.Fatal(err)
	}
	if curve != "secp256k1" {
		log.Fatal(fmt.Errorf("must use secp256k1 private key, but found %s", curve))
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatal("crypto.ToECDSA")
	}
	address = crypto.PubkeyToAddress(privateKey.PublicKey)

	return address, privateKeyBytes
}

// GetAddressFromGenerating ...
func GetAddressFromGenerating() (address common.Address, privateKeyBytes []byte) {

	// 生成随机私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// 转换私钥为字节序列
	privateKeyBytes = crypto.FromECDSA(privateKey)

	// 转换公钥为地址
	address = crypto.PubkeyToAddress(privateKey.PublicKey)

	return address, privateKeyBytes
}
