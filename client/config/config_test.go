package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var testConfig1 = []byte(`

# DeFaaS 配置文件

title = "DeFaaS Config"

[contracts]
FaasToken = "0xbb8990586d4de6634f97a094591644fae047502c"
Market = "0xc79a7fd8448609d2f6de564a3aad153eace6ae97"
WitnessPool = "0xeee6ea03067cee5835ab898417b80484d4fd8a67"

`)

func checkTestConfig1(dfc *DeFaaSConfig) bool {

	if !strings.EqualFold(dfc.FaaSTokenContractAddress.Hex(), "0xbb8990586d4de6634f97a094591644fae047502c") {
		return false
	}
	if !strings.EqualFold(dfc.MarketContractAddress.Hex(), "0xc79a7fd8448609d2f6de564a3aad153eace6ae97") {
		return false
	}
	if !strings.EqualFold(dfc.WitnessPoolContractAddress.Hex(), "0xeee6ea03067cee5835ab898417b80484d4fd8a67") {
		return false
	}

	return true
}

func TestParseConfig(t *testing.T) {

	dfc, err := ParseConfig(bytes.NewBuffer(testConfig1))
	if err != nil {
		t.Fatal(err)
	}

	if !checkTestConfig1(dfc) {
		t.Fatal()
	}
}

func TestParseConfigFile(t *testing.T) {

	tmpfile, err := ioutil.TempFile("", "test-defaas-config")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(testConfig1); err != nil {
		t.Fatal(err)
	}

	dfc, err := ParseConfigFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if !checkTestConfig1(dfc) {
		t.Fatal()
	}
}

func checkTestConfig2(dfc *DeFaaSConfig) bool {

	if !strings.EqualFold(dfc.FaaSTokenContractAddress.Hex(), "0x0000000000000000000000000000000000000001") {
		return false
	}
	if !strings.EqualFold(dfc.MarketContractAddress.Hex(), "0x0000000000000000000000000000000000000002") {
		return false
	}
	if !strings.EqualFold(dfc.WitnessPoolContractAddress.Hex(), "0x0000000000000000000000000000000000000003") {
		return false
	}

	return true
}

func TestWriteConfig(t *testing.T) {

	tmpfile, err := ioutil.TempFile("", "test-defaas-config")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(testConfig1); err != nil {
		t.Fatal(err)
	}

	if err := WriteConfig(tmpfile.Name(),
		common.HexToAddress("0x0000000000000000000000000000000000000001"),
		common.HexToAddress("0x0000000000000000000000000000000000000002"),
		common.HexToAddress("0x0000000000000000000000000000000000000003")); err != nil {
		t.Fatal(err)
	}

	dfc, err := ParseConfigFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if !checkTestConfig2(dfc) {
		t.Fatal()
	}
}
