package testconfig

import (
	"log"
	"path"

	devutils "defaas/dev-cmd/utils"
)

type TestConfig struct {
	TestDeFaaSConfigFilePath string
	TestKeyStoreFilePath     string
	TestKeyStorePassword     string
}

func GetTestConfig() *TestConfig {

	tcfg := &TestConfig{}

	workDir := "/home/dds/kitchen/defaas"

	tcfg.TestDeFaaSConfigFilePath = path.Join(workDir, "defaas-config.toml")
	testKeyStoreDir := path.Join(workDir, "private-chain/data-0/keystore")
	names, err := devutils.ReadDirNames(testKeyStoreDir)
	if err != nil {
		log.Fatal(err)
	}

	tcfg.TestKeyStoreFilePath = path.Join(testKeyStoreDir, names[0])
	tcfg.TestKeyStorePassword = ""

	return tcfg
}