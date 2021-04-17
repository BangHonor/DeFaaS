package wintess

import (
	basic "defaas/client/basic"
	"testing"
)

func getTestWitnessClientFromFile() (*WitnessClient, error) {

	tcfg := basic.GetTestConfig()

	return NewWitnessClientWithFile(
		tcfg.TestDeFaaSConfigFilePath,
		tcfg.TestKeyStoreFilePath,
		tcfg.TestKeyStorePassword)
}

func TestNewWitnessClient(t *testing.T) {

	client, err := getTestWitnessClientFromFile()
	if err != nil {
		t.Fatal(err)
	}

	_ = client
}

func TestLogin(t *testing.T) {

	client, err := getTestWitnessClientFromFile()
	if err != nil {
		t.Fatal(err)
	}

	if err := client.Login(); err != nil {
		t.Fatal(err)
	}
}

func TestLogout(t *testing.T) {

	client, err := getTestWitnessClientFromFile()
	if err != nil {
		t.Fatal(err)
	}

	if err := client.Logout(); err != nil {
		t.Fatal(err)
	}
}

func TestLoginLogout(t *testing.T) {

	client, err := getTestWitnessClientFromFile()
	if err != nil {
		t.Fatal(err)
	}

	if err := client.Login(); err != nil {
		t.Fatal(err)
	}

	if err := client.Logout(); err != nil {
		t.Fatal(err)
	}
}

func TestTurnOnTurnOff(t *testing.T) {

	client, err := getTestWitnessClientFromFile()
	if err != nil {
		t.Fatal(err)
	}

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
