package customer

import (
	"defaas/tests/testconfig"
	"log"
	"testing"
)

func getTestCustomerClientFromFile() *CustomerClient {

	tcfg := testconfig.GetTestConfig()

	client, err := NewCustomerClientWithFile(
		tcfg.TestDeFaaSConfigFilePath,
		tcfg.TestKeyStoreFilePath,
		tcfg.TestKeyStorePassword)

	if err != nil {
		log.Fatal(err)
	}

	return client
}
func TestNewCustomerClient(t *testing.T) {

	client := getTestCustomerClientFromFile()
	_ = client
}
