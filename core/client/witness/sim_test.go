package witness

import (
	"defaas/core/suite"
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func TestMonitoring1(t *testing.T) {

	sim := suite.NewSim()
	blockchain := sim.Blockchain
	witnesses := []*WitnessClient{}

	for i := 0; i < len(sim.Keys); i++ {

		witness := &WitnessClient{}

		witness.Key = sim.Keys[i]
		witness.DeFaaSConfig = nil
		witness.ETHClient = nil

		auth, err := bind.NewKeyedTransactorWithChainID(sim.Keys[i].PrivateKey, blockchain.Blockchain().Config().ChainID)
		if err != nil {
			log.Fatal(err)
		}
		witness.Suite, err = suite.NewSuite(blockchain, auth, sim.Suite.FaaSTokenAddress, sim.Suite.MarketAddress, sim.Suite.WitnessPoolAddress)
		if err != nil {
			log.Fatal(err)
		}

		witness.stopRunningTrigger = make(chan struct{})

		witnesses = append(witnesses, witness)
	}

	if onlines, err := witnesses[0].WitnessPool.NumOnlineWitness(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("onlines [%v]\n", onlines)
	}

	for i := 0; i < len(sim.Keys); i++ {

		if err := witnesses[i].Login(); err != nil {
			log.Fatal(err)
		}
	}
	blockchain.Commit()

	if onlines, err := witnesses[0].WitnessPool.NumOnlineWitness(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("onlines [%v]\n", onlines)
	}

	for i := 0; i < len(sim.Keys); i++ {

		if err := witnesses[i].TurnOn(); err != nil {
			log.Fatal(err)
		}
	}
	blockchain.Commit()

	if onlines, err := witnesses[0].WitnessPool.NumOnlineWitness(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("onlines [%v]\n", onlines)
	}

	for i := 0; i < len(sim.Keys); i++ {

		if err := witnesses[i].TurnOff(); err != nil {
			log.Fatal(err)
		}
	}
	blockchain.Commit()

	if onlines, err := witnesses[0].WitnessPool.NumOnlineWitness(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("onlines [%v]\n", onlines)
	}

	{
		curBlockNum := blockchain.Blockchain().CurrentBlock().Number()

		blockNeed, err := witnesses[0].WitnessPool.StdBlockNeed()
		if err != nil {
			log.Fatal(err)
		}
		for i := big.NewInt(0); i.Cmp(blockNeed) == -1; i = i.Add(i, big.NewInt(1)) {
			blockchain.Commit()
		}

		slaID := big.NewInt(233)
		monitoringDuration := big.NewInt(1)

		log.Println("newSLA ...")
		newslaTx, err := witnesses[0].WitnessPool.NewSLA(curBlockNum, slaID, "", monitoringDuration)
		if err != nil {
			log.Fatal(err)
		}
		_ = newslaTx
		blockchain.Commit()
		log.Println("newSLA done")

	}

	time.Sleep(3 * time.Second)
}

func TestMonitoring2(t *testing.T) {

	sim := suite.NewSim()
	blockchain := sim.Blockchain
	witnesses := []*WitnessClient{}

	for i := 0; i < len(sim.Keys); i++ {

		witness := &WitnessClient{}

		witness.Key = sim.Keys[i]
		witness.DeFaaSConfig = nil
		witness.ETHClient = nil

		auth, err := bind.NewKeyedTransactorWithChainID(sim.Keys[i].PrivateKey, blockchain.Blockchain().Config().ChainID)
		if err != nil {
			log.Fatal(err)
		}
		witness.Suite, err = suite.NewSuite(blockchain, auth, sim.Suite.FaaSTokenAddress, sim.Suite.MarketAddress, sim.Suite.WitnessPoolAddress)
		if err != nil {
			log.Fatal(err)
		}

		witness.stopRunningTrigger = make(chan struct{})

		witnesses = append(witnesses, witness)

		if err := witnesses[i].Login(); err != nil {
			log.Fatal(err)
		}
		blockchain.Commit()

		if err := witnesses[i].TurnOn(); err != nil {
			log.Fatal(err)
		}
		blockchain.Commit()
	}

	{
		curBlockNum := blockchain.Blockchain().CurrentBlock().Number()

		blockNeed, err := witnesses[0].WitnessPool.StdBlockNeed()
		if err != nil {
			log.Fatal(err)
		}
		for i := big.NewInt(0); i.Cmp(blockNeed) == -1; i = i.Add(i, big.NewInt(1)) {
			blockchain.Commit()
		}

		slaID := big.NewInt(233)
		monitoringDuration := big.NewInt(1)

		log.Println("newSLA ...")
		newslaTx, err := witnesses[0].WitnessPool.NewSLA(curBlockNum, slaID, "", monitoringDuration)
		if err != nil {
			log.Fatal(err)
		}
		_ = newslaTx
		blockchain.Commit()
		log.Println("newSLA done")

		for i := 0; i < len(witnesses); i++ {
			if err := witnesses[i].Report(slaID, true); err != nil {
				log.Println(err)
			}
		}

		log.Println("check ...")
		checkslaTx, err := witnesses[0].WitnessPool.CheckSLA(slaID)
		if err != nil {
			log.Fatal(err)
		}
		_ = checkslaTx
		blockchain.Commit()
		log.Println("check done")

		time.Sleep(3 * time.Second)

		if isViolated, err := witnesses[0].WitnessPool.IsViolatedSLA(slaID); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("isViolated [%v]\n", isViolated)
		}
	}

}
