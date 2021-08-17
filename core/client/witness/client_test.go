package witness

import (
	"context"
	"defaas/core/helper"
	"defaas/tests/testconfig"
	"log"
	"math/big"
	"testing"
	"time"
)

func getTestWitnessClientFromFile() *WitnessClient {

	tcfg := testconfig.GetTestConfig()

	client, err := NewWitnessClientWithFile(
		tcfg.TestDeFaaSConfigFilePath,
		tcfg.TestKeyStoreFilePath,
		tcfg.TestKeyStorePassword)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func TestNewWitnessClient(t *testing.T) {

	client := getTestWitnessClientFromFile()
	_ = client
}

func TestLogin(t *testing.T) {

	client := getTestWitnessClientFromFile()
	client.FaaSTokenEventSub(context.TODO())

	if err := client.Login(); err != nil {
		t.Fatal(err)
	}
}

func TestLogout(t *testing.T) {

	client := getTestWitnessClientFromFile()
	client.FaaSTokenEventSub(context.TODO())

	if err := client.Logout(); err != nil {
		t.Fatal(err)
	}
}

func TestLoginLogout(t *testing.T) {

	client := getTestWitnessClientFromFile()

	client.FaaSTokenEventSub(context.TODO())

	if err := client.Login(); err != nil {
		t.Fatal(err)
	}

	if err := client.Logout(); err != nil {
		t.Fatal(err)
	}
}

func TestTurnOnTurnOff(t *testing.T) {

	client := getTestWitnessClientFromFile()

	if err := client.Login(); err != nil {
		t.Fatal(err)
	}

	if err := client.TurnOn(); err != nil {
		t.Fatal(err)
	}

	if err := client.TurnOff(); err != nil {
		t.Fatal(err)
	}

	if err := client.Logout(); err != nil {
		t.Fatal(err)
	}
}

func getWitnesses() []*WitnessClient {

	ksPaths := testconfig.GetTestKeyStorePaths()

	tcfg := testconfig.GetTestConfig()

	witnesses := []*WitnessClient{}

	for i := 0; i < len(ksPaths); i++ {

		log.Printf("[%v/%v]\n", i, len(ksPaths))

		witness, err := NewWitnessClientWithFile(tcfg.TestDeFaaSConfigFilePath, ksPaths[i], "")
		if err != nil {
			log.Fatal(err)
		}

		witnesses = append(witnesses, witness)
	}

	return witnesses
}

func TestManyLogin(t *testing.T) {

	witnesses := getWitnesses()

	for i := 0; i < len(witnesses); i++ {

		log.Printf("[%v/%v]\n", i, len(witnesses))

		registered, err := witnesses[i].WitnessPool.IsWitnessRegistered(witnesses[i].Key.Address)
		if err != nil {
			log.Fatal(err)
		}

		if !registered {
			if err := witnesses[i].Login(); err != nil {
				log.Fatal(err)
			}
		}
	}

	if onlines, err := witnesses[0].WitnessPool.NumOnlineWitness(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("onlines [%v]\n", onlines)
	}
}

func TestManyLogout(t *testing.T) {

	witnesses := getWitnesses()

	for i := 0; i < len(witnesses); i++ {

		log.Printf("[%v/%v]\n", i, len(witnesses))

		registered, err := witnesses[i].WitnessPool.IsWitnessRegistered(witnesses[i].Key.Address)
		if err != nil {
			log.Fatal(err)
		}

		if registered {
			if err := witnesses[i].Logout(); err != nil {
				log.Fatal(err)
			}
		}
	}

	if onlines, err := witnesses[0].WitnessPool.NumOnlineWitness(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("onlines [%v]\n", onlines)
	}
}

func TestManyTurnOn(t *testing.T) {

	witnesses := getWitnesses()

	for i := 0; i < len(witnesses); i++ {

		log.Printf("[%v/%v]\n", i, len(witnesses))

		if err := witnesses[i].TurnOn(); err != nil {
			t.Log(err)
		}
	}

	if onlines, err := witnesses[0].WitnessPool.NumOnlineWitness(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("onlines [%v]\n", onlines)
	}
}

func TestManyTurnOff(t *testing.T) {

	witnesses := getWitnesses()

	for i := 0; i < len(witnesses); i++ {

		log.Printf("[%v/%v]\n", i, len(witnesses))

		if err := witnesses[i].TurnOff(); err != nil {
			t.Log(err)
		}
	}

	if onlines, err := witnesses[0].WitnessPool.NumOnlineWitness(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("onlines [%v]\n", onlines)
	}
}

func TestManyLoginTurnOn(t *testing.T) {

	witnesses := getWitnesses()

	for i := 0; i < len(witnesses); i++ {

		log.Printf("[%v/%v]\n", i, len(witnesses))

		registered, err := witnesses[i].WitnessPool.IsWitnessRegistered(witnesses[i].Key.Address)
		if err != nil {
			log.Fatal(err)
		}

		if !registered {
			if err := witnesses[i].Login(); err != nil {
				log.Fatal(err)
			}
		}

		if err := witnesses[i].TurnOn(); err != nil {
			t.Log(err)
		}
	}

	if onlines, err := witnesses[0].WitnessPool.NumOnlineWitness(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("onlines [%v]\n", onlines)
	}
}

func TestWitnessGame(t *testing.T) {

	witnesses := getWitnesses()

	//
	for i := 0; i < len(witnesses); i++ {
		log.Printf("[%v/%v]\n", i, len(witnesses))
		witnesses[i].StartWatching()
	}

	//
	if onlines, err := witnesses[0].WitnessPool.NumOnlineWitness(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("onlines [%v]\n", onlines)
	}

	//
	curBlockNum := helper.CurrentBlockNumber(witnesses[0].ETHClient)

	blockNeed, err := witnesses[0].WitnessPool.StdBlockNeed()
	if err != nil {
		log.Fatal(err)
	}
	helper.WaitMinedBlocksBySubscription(witnesses[0].ETHClient, int(blockNeed.Int64()))

	slaID := big.NewInt(time.Now().Unix())
	monitoringDuration := big.NewInt(1)

	log.Println("newSLA ...")
	newslaTx, err := witnesses[0].WitnessPool.NewSLA(curBlockNum, slaID, "", monitoringDuration)
	if err != nil {
		log.Fatal(err)
	}
	witnesses[0].ConfirmTxByPolling(newslaTx.Hash(), helper.NumBlockToWaitRecommended)
	log.Println("newSLA done")

	if onlines, err := witnesses[0].WitnessPool.NumOnlineWitness(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("onlines [%v]\n", onlines)
	}

	if _, err := witnesses[0].WitnessPool.IsViolatedSLA(slaID); err == nil {
		log.Fatal()
	}

	log.Println("checkSLA ...")
	checkslaTx, err := witnesses[0].WitnessPool.CheckSLA(slaID)
	if err != nil {
		log.Fatal(err)
	}
	witnesses[0].ConfirmTxBySubscription(checkslaTx.Hash(), helper.NumBlockToWaitRecommended)
	log.Println("checkSLA done")

	time.Sleep(15 * time.Second)

	if isViolated, err := witnesses[0].WitnessPool.IsViolatedSLA(slaID); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("isViolated [%v]\n", isViolated)
	}

	for i := 0; i < len(witnesses); i++ {
		log.Printf("[%v/%v]\n", i, len(witnesses))
		witnesses[i].StoptWatching()
	}

	//
	if onlines, err := witnesses[0].WitnessPool.NumOnlineWitness(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("onlines [%v]\n", onlines)
	}
}
