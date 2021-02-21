package utils

import (
	"log"

	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// GetAddressFromConfig ...
func GetAddressFromConfig(config *conf.Config) common.Address {
	privateKey, err := crypto.ToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal("parse error")
	}
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}
