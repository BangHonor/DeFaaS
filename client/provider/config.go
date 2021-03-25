package provider

import (
	"io"

	"github.com/spf13/viper"
)

type ProviderConfig struct {
	Adapter     string
	ServerAddr  string
	ServerEntry string
}

const (
	adapterKey     = "deploypath.adapter"
	serverAddrKey  = "deploypath.serveraddr"
	serverEntryKey = "deploypath.serverentry"
)

func ParseConfigFile(providerConfigFilePath string) (*ProviderConfig, error) {
	pc := &ProviderConfig{}

	// TODO

	return pc, nil
}

func ParseConfig(in io.Reader) (*ProviderConfig, error) {

	pc := &ProviderConfig{}

	viper.SetConfigType("toml")
	if err := viper.ReadConfig(in); err != nil {
		return nil, err
	}

	{
		pc.Adapter = viper.GetString(adapterKey)
		pc.ServerAddr = viper.GetString(serverAddrKey)
		pc.ServerEntry = viper.GetString(serverEntryKey)
	}

	return pc, nil
}
