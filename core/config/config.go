package config

import (
	"errors"
	"io"
	"os"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

const (
	// key in config file
	faastokenKey   = "contracts.faastoken"
	marketKey      = "contracts.market"
	witnesspoolKey = "contracts.witnesspool"
	HttpURLsKey    = "network.httpurls"
	wsURLsKey      = "network.wsurls"
)

// 区块链节点地址
type ChainNodeURL struct {
	HttpURL string
	WsURL   string
}

// DeFaaSConfig ...
type DeFaaSConfig struct {
	HttpURLs                   []string
	WsURLs                     []string
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
		return nil, err
	}
	defer f.Close()

	return ParseConfig(f)
}

// ParseConfig ...
func ParseConfig(in io.Reader) (*DeFaaSConfig, error) {
	fc := &DeFaaSConfig{}

	viper.SetConfigType("toml")
	if err := viper.ReadConfig(in); err != nil {
		return nil, err
	}

	{
		// 检查以太坊地址是否有效
		re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

		_FaaSTokenContractHex := viper.GetString(faastokenKey)
		_MarketContractHex := viper.GetString(marketKey)
		_WitnessPoolContractHex := viper.GetString(witnesspoolKey)

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
	}

	{
		fc.HttpURLs = viper.GetStringSlice(HttpURLsKey)
		fc.WsURLs = viper.GetStringSlice(wsURLsKey)
	}

	return fc, nil
}

// WriteContractAddress ...
// 修改配置文件中的合约地址, 不改变配置文件的其他字段
func WriteContractAddress(configFilePath string, faastokenContractAddress, marketContractAddress, witnesspoolContractAddress common.Address) error {

	// https://github.com/spf13/viper#writing-config-files

	f, err := os.OpenFile(configFilePath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	viper.SetConfigType("toml")
	viper.ReadConfig(f)

	viper.Set(faastokenKey, faastokenContractAddress.Hex())
	viper.Set(marketKey, marketContractAddress.Hex())
	viper.Set(witnesspoolKey, witnesspoolContractAddress.Hex())

	if err := viper.WriteConfigAs(configFilePath); err != nil {
		return err
	}

	return nil
}
