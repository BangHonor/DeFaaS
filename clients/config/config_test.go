package config

import (
	"bytes"
	"testing"
)

var testConfig = []byte(`
# DeFaaS 配置文件

title = "DeFaaS Config"

[contracts]
FaasToken = "0xbb8990586d4de6634f97a094591644fae047502c"
Market = "0xc79a7fd8448609d2f6de564a3aad153eace6ae97"
WitnessPool = "0xeee6ea03067cee5835ab898417b80484d4fd8a67"

`)

func TestParseConfig(t *testing.T) {
	dfc, err := ParseConfig(bytes.NewBuffer(testConfig))
	if err != nil {
		t.Fatal()
	}
	t.Log(dfc.FaaSTokenContractAddress.Hex())
	t.Log(dfc.MarketContractAddress.Hex())
	t.Log(dfc.WitnessPoolContractAddress.Hex())
}
