package config

type ServerConfig struct {
	Port       string `mapstructure:"port"`
	Host       string `mapstructure:"host"`
	Encryption int    `mapstructure:"encryption-cost"`
}
