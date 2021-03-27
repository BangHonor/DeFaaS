package provider

import (
	"defaas/core/data"
	"io"
	"os"

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

	f, err := os.OpenFile(providerConfigFilePath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ParseConfig(f)
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

func GetDeployPathFromProviderConfig(pc *ProviderConfig) string {

	return data.GetDeployPath(pc.Adapter, pc.ServerAddr, pc.ServerEntry)

}
