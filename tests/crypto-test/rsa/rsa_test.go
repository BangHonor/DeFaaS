package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestGenerateRSAKey(t *testing.T) {

	assert := assert.New(t)

	pk, err := GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(ConstKeyBytes, pk.Size())
}

func TestEncodeDecodePublicKey(t *testing.T) {

	assert := assert.New(t)

	pk, err := GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	encoded, err := EncodePublicKey(pk.Public())
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(encoded))

	pub, err := DecodePublicKey(encoded)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(pk.Public(), pub)

	{
		_encoded, err := EncodePublicKey(pub)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(encoded, _encoded)

	}

}

func TestEncryptDecrypt(t *testing.T) {

	assert := assert.New(t)

	pk, err := GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	raw := []byte(`Hello`)

	encrypted, err := PublicKeyEncrypt(pk.Public(), raw)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(raw))
	t.Log(len(encrypted))

	decrypted, err := PrivateKeyDecrypt(pk, encrypted)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(raw, decrypted)
}

func TestEncodeDecodePublicKeyEncryptDecrypt(t *testing.T) {

	assert := assert.New(t)

	pk, err := GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	raw := []byte(`Hello`)

	encoded, err := EncodePublicKey(pk.Public())
	if err != nil {
		t.Fatal(err)
	}

	pub, err := DecodePublicKey(encoded)
	if err != nil {
		t.Fatal(err)
	}

	encrypted, err := PublicKeyEncrypt(pub, raw)
	if err != nil {
		t.Fatal(err)
	}

	decrypted, err := PrivateKeyDecrypt(pk, encrypted)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(raw, decrypted)
}

func sizedBytes(n int) []byte {
	bs := make([]byte, n)
	return bs
}

func TestSizedInputEncrypt(t *testing.T) {

	assert := assert.New(t)

	pk, err := GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	raws := func() [][]byte {
		raws := [][]byte{}
		for i := 0; i <= ConstValidRawDataBytes; i++ {
			raws = append(raws, sizedBytes(i))
		}
		return raws
	}()

	for _, raw := range raws {
		_, err := PublicKeyEncrypt(pk.Public(), raw)
		assert.Nil(err)
	}

	{
		_, err := PublicKeyEncrypt(pk.Public(), sizedBytes(ConstValidRawDataBytes+1))
		assert.NotNil(err)
	}
}
