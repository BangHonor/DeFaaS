package testconfig

import (
	"fmt"
	"testing"
)

func TestGetTestConfig(t *testing.T) {
	tcfg := GetTestConfig()
	fmt.Println(tcfg.TestDeFaaSConfigFilePath)
	fmt.Println(tcfg.TestKeyStoreFilePath)
	fmt.Println(tcfg.TestKeyStorePassword)
}

func TestGenerateTestKeyStore(t *testing.T) {
	GenerateTestKeyStore()
}
