package suite

import (
	"log"
	"testing"
)

func TestSuitStdWitnessDepoist(t *testing.T) {

	sim := NewSimSuite()

	deposit, err := sim.WitnessPool.StdWitnessDepoist()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("deposit", deposit)

	log.Println(sim.Suite)
}

func simtest() {
	sim := NewSimSuite()

	customer := sim.Others[0]
	provider := sim.Others[1]

	log.Println(customer.From)
	log.Println(provider.From)
}

func TestSimTest(t *testing.T) {
	simtest()
}
