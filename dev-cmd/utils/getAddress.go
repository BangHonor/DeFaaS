package utils

// 产生地址

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

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
