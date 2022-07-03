package config

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}
