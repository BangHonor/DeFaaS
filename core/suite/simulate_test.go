package suite

import (
	"log"
	"testing"
)

func TestSuitStdWitnessDepoist(t *testing.T) {

	sim := NewSim()

	deposit, err := sim.WitnessPool.StdWitnessDepoist()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("deposit", deposit)

	log.Println(sim.Suite)
}

func TestSim1(t *testing.T) {

	sim := NewSim()

	customer := sim.Keys[0].Address
	provider := sim.Keys[1].Address

	t.Log(customer)
	t.Log(provider)
}
