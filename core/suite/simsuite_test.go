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

	customerAuth := sim.Others[0]
	providerAuth := sim.Others[1]

	log.Println(customerAuth.From)
	log.Println(providerAuth.From)

	go func() {

	}()

	go func() {

	}()

	for i := 0; i < 10; i++ {
		go func() {

		}()
	}
}

func TestSimTest(t *testing.T) {
	simtest()
}
