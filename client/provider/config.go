package provider

type ProviderConfig struct {
	DeployPath string
}

func ParseConfigFile(providerConfigFilePath string) (*ProviderConfig, error) {
	pc := &ProviderConfig{}

	// TODO

	return pc, nil
}
