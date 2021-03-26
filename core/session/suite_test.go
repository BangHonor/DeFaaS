package session

import "testing"

func TestSuitStdWitnessDepoist(t *testing.T) {

	suit := simSingleCase()

	session := suit.witnesspool

	deposit, err := session.StdWitnessDepoist()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("deposit", deposit)
}
