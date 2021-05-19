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
}

func TestManyLogout(t *testing.T) {

	witnesses := getWitnesses()

	for i := 0; i < len(witnesses); i++ {

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
}

func TestManyTurnOn(t *testing.T) {

	witnesses := getWitnesses()

	for i := 0; i < len(witnesses); i++ {

		if err := witnesses[i].TurnOn(); err != nil {
			t.Log(err)
		}
	}
}

func TestManyTurnOff(t *testing.T) {

	witnesses := getWitnesses()

	for i := 0; i < len(witnesses); i++ {

		if err := witnesses[i].TurnOff(); err != nil {
			t.Log(err)
		}
	}
}

func TestWitnessGame(t *testing.T) {
	WitnessGame()
}

func WitnessGame() {

	witnesses := getWitnesses()

	for i := 0; i < len(witnesses); i++ {
		witnesses[i].StartWatching()
	}

	curBlockNum := helper.CurrentBlockNumber(witnesses[0].ETHClient)

	blockNeed, err := witnesses[0].WitnessPool.StdBlockNeed()
	if err != nil {
		log.Fatal(err)
	}
	helper.WaitMinedBlocksBySubscription(witnesses[0].ETHClient, int(blockNeed.Int64()))

	slaID := big.NewInt(233)
	monitoringDuration := big.NewInt(1)

	log.Println("newSLA ...")
	newslaTx, err := witnesses[0].WitnessPool.NewSLA(curBlockNum, slaID, "", monitoringDuration)
	if err != nil {
		log.Fatal(err)
	}
	witnesses[0].ConfirmTxByPolling(newslaTx.Hash(), helper.NumBlockToWaitRecommended)
	log.Println("newSLA done")

	log.Println("checkSLA ...")
	checkslaTx, err := witnesses[0].WitnessPool.CheckSLA(slaID)
	if err != nil {
		log.Fatal(err)
	}
	witnesses[0].ConfirmTxBySubscription(checkslaTx.Hash(), helper.NumBlockToWaitRecommended)
	log.Println("checkSLA done")

	time.Sleep(2 * time.Second)

	if isViolated, err := witnesses[0].WitnessPool.IsViolatedSLA(slaID); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("isViolated [%v]\n", isViolated)
	}

	for i := 0; i < len(witnesses); i++ {
		if err := witnesses[i].Logout(); err != nil {
			log.Fatal(err)
		}
	}
}
