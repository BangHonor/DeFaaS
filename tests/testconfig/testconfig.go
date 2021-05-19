package testconfig

import (
	"log"
	"path"

	devutils "defaas/dev-cmd/utils"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

var (
	workDir = "/home/kitchen/defaas"
)

type TestConfig struct {
	TestDeFaaSConfigFilePath string
	TestKeyStoreFilePath     string
	TestKeyStorePassword     string
}

func GetTestConfig() *TestConfig {

	tcfg := &TestConfig{}

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

func GenerateTestKeyStore() {

	ksDirPath := path.Join(workDir, "tests/testconfig/keystore")

	ks := keystore.NewKeyStore(ksDirPath, keystore.StandardScryptN, keystore.StandardScryptP)
	password := ""

	for i := 0; i < 10; i++ {
		account, err := ks.NewAccount(password)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(account.Address)
	}
}

func GetTestKeyStorePaths() []string {

	ksDirPath := path.Join(workDir, "tests/testconfig/keystore")
	names, err := devutils.ReadDirNames(ksDirPath)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(names); i++ {
		names[i] = path.Join(ksDirPath, names[i])
	}

	return names
}
