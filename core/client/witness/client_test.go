package witness

import (
	"context"
	"defaas/tests/testconfig"
	"log"
	"testing"
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
