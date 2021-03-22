package sha256

import (
	"crypto/sha256"
	"math/big"
	"testing"
)

// https://www.tutorialspoint.com/solidity/solidity_cryptographic_functions.htm
// https://pkg.go.dev/crypto/sha256#Sum256

func TestSha256(t *testing.T) {

	input := []byte("hello world")
	var hashValue [32]byte = [32]byte{}

	hashValue = sha256.Sum256(input)

	t.Log(hashValue)

	t.Log(big.NewInt(0).SetBytes(hashValue[:]).Text(0x10))
}
