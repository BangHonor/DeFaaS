package provider

type ProviderConfig struct {
	Adapter     string
	ServerAddr  string
	ServerEntry string
}

func ParseConfigFile(providerConfigFilePath string) (*ProviderConfig, error) {
	pc := &ProviderConfig{}

	// TODO

	return pc, nil
}
