package config

import (
	"io"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

// https://github.com/spf13/viper

// DeFaaSConfig ...
type DeFaaSConfig struct {
	FaaSTokenContractAddress   common.Address
	MarketContractAddress      common.Address
	WitnessPoolContractAddress common.Address
}

// ParseConfigFile ...
// usage:
//     dfc, _ := ParseConfigFile("defaas-config.toml")
func ParseConfigFile(configFilePath string) (*DeFaaSConfig, error) {

	f, err := os.OpenFile(configFilePath, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(f)
	}
	defer f.Close()

	return ParseConfig(f)
}

// ParseConfig ...
func ParseConfig(in io.Reader) (*DeFaaSConfig, error) {
	fc := &DeFaaSConfig{}

	viper.SetConfigType("toml")
	viper.ReadConfig(in)

	fc.FaaSTokenContractAddress = common.HexToAddress(viper.GetString("contracts.FaaSToken"))
	fc.MarketContractAddress = common.HexToAddress(viper.GetString("contracts.Market"))
	fc.WitnessPoolContractAddress = common.HexToAddress(viper.GetString("contracts.WitnessPool"))

	return fc, nil
}
