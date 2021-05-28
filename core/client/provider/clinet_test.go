package provider

import (
	"defaas/tests/testconfig"
	"log"
	"testing"
)

func getTestProviderClientFromFile() *ProviderClient {

	tcfg := testconfig.GetTestConfig()

	client, err := NewProviderClientWithFile(
		tcfg.TestDeFaaSConfigFilePath,
		tcfg.TestKeyStoreFilePath,
		tcfg.TestKeyStorePassword)

	if err != nil {
		log.Fatal(err)
	}

	return client
}
func TestNewProviderClient(t *testing.T) {

	client := getTestProviderClientFromFile()
	_ = client
}
