package config

import (
	"errors"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

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

	// 检查以太坊地址是否有效
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	_FaaSTokenContractHex := viper.GetString("contracts.FaaSToken")
	_MarketContractHex := viper.GetString("contracts.Market")
	_WitnessPoolContractHex := viper.GetString("contracts.WitnessPool")

	if !re.MatchString(_FaaSTokenContractHex) {
		return nil, errors.New("invalid FaaSTokenContractAddress hex")
	}
	if !re.MatchString(_MarketContractHex) {
		return nil, errors.New("invalid MarketContractAddress hex")
	}
	if !re.MatchString(_WitnessPoolContractHex) {
		return nil, errors.New("invalid WitnessPoolContractAddress hex")
	}

	fc.FaaSTokenContractAddress = common.HexToAddress(_FaaSTokenContractHex)
	fc.MarketContractAddress = common.HexToAddress(_MarketContractHex)
	fc.WitnessPoolContractAddress = common.HexToAddress(_WitnessPoolContractHex)

	return fc, nil
}
