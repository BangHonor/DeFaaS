package suite

import "testing"

func TestSuitStdWitnessDepoist(t *testing.T) {

	suite := SingleSimSuite()

	deposit, err := suite.WitnessPool.StdWitnessDepoist()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("deposit", deposit)
}
