package session

import (
	"testing"
)

func TestStdWitnessDepoist(t *testing.T) {

	auth, blockchain, _, session := newSimWitnessPool()
	_ = auth
	_ = blockchain
	_ = session

	deposit, err := session.StdWitnessDepoist()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("deposit", deposit)
}
