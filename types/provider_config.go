package types

type ProviderConfig struct {
	VaultDefaultPolicy    string
	VaultSecretPathPrefix string
	Datacenter            string
	ConsulAddress         string
	ConsulDNSEnabled      bool
}
