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

func TestGetTestKeyStorePaths(t *testing.T) {

	paths := GetTestKeyStorePaths()
	for _, path := range paths {
		t.Log(path)
	}
}
