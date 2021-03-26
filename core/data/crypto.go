package data

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	mathrand "math/rand"
	"time"
)

type PublicKey = crypto.PublicKey

type PrivateKey interface {
	Public() PublicKey
	Size() int
}

const (
	ConstKeyBytes          = 32
	ConstValidRawDataBytes = 21 // depends on `ConstKeyBytes`
)

func GeneratePrivateKey() (PrivateKey, error) {

	privateKey, err := rsa.GenerateKey(rand.Reader, ConstKeyBytes*8)

	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func EncodePublicKey(pub PublicKey) ([]byte, error) {

	encoded, err := x509.MarshalPKIXPublicKey(pub)

	if err != nil {
		return nil, err
	}

	return encoded, nil
}

func DecodePublicKey(encoded []byte) (PublicKey, error) {

	pubInterface, err := x509.ParsePKIXPublicKey([]byte(encoded))

	if err != nil {
		return nil, err
	}

	return pubInterface.(*rsa.PublicKey), nil
}

func PublicKeyEncrypt(publicKey PublicKey, rawData []byte) ([]byte, error) {

	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), rawData)

	if err != nil {
		return nil, err
	}

	return encryptedData, nil
}

func PrivateKeyDecrypt(privateKey PrivateKey, encryptedData []byte) ([]byte, error) {

	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), encryptedData)

	if err != nil {
		return nil, err
	}

	return decryptedData, nil

}

// --------------------------------------------------------------------------

// GenerateFulfillSecretKey returns a random 32-length byte array
func GenerateFulfillSecretKey() [32]byte {

	// Note that this implementation is very problematic !!!

	var sk [32]byte

	mathrand.Seed(time.Now().Unix())
	copy(sk[:32:32], []byte(fmt.Sprintf("%0*v", 32, mathrand.Int())))

	return sk
}

func EncryptAccessKey(accessKey, publicKey string) (string, error) {

	pub, err := DecodePublicKey([]byte(publicKey))
	if err != nil {
		return "", nil
	}

	encryptedAccessKey, err := PublicKeyEncrypt(pub, []byte(accessKey))
	if err != nil {
		return "", nil
	}

	return string(encryptedAccessKey), nil
}

func DecryptAccessKey(encryptedAccessKey string, privateKey PrivateKey) (string, error) {

	accessKey, err := PrivateKeyDecrypt(privateKey, []byte(encryptedAccessKey))

	if err != nil {
		return "", err
	}

	return string(accessKey), nil
}
