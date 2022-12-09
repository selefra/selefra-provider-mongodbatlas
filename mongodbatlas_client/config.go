package mongodbatlas_client

type Configs struct {
	Providers []Config `yaml:"providers"  mapstructure:"providers"`
}

type Config struct {
	PublicKey  string `yaml:"public_key,omitempty" mapstructure:"public_key"`
	PrivateKey string `yaml:"private_key,omitempty" mapstructure:"private_key"`
}
