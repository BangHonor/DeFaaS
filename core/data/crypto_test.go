package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {

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

func TestGenerateFulfillSecretKey(t *testing.T) {

	sk := GenerateFulfillSecretKey()

	t.Log(string(sk[:]))
}
