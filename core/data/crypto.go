package data

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

func GenerateRSAPrivateKey() *rsa.PrivateKey {

	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)

	if err != nil {
		log.Fatal(err)
	}

	return privateKey
}

// func RSAEncrypt(data, rsaPublicKey string) string {

// 	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)

// 	return ""
// }
