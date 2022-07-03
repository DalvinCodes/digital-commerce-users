package config

import (
	"github.com/spf13/viper"
	"log"
)

type Configurations struct {
	Server   ServerConfig `mapstructure:"server"`
	Postgres UsersDB      `mapstructure:"postgres"`
}

const FileName = "api-user-v1.yaml"

func LoadConfigs() *Configurations {
	v, err := initConfig(FileName)
	if err != nil {
		log.Fatal(err)
	}
	configs, err := parseConfig(v)
	if err != nil {
		log.Fatal(err)
	}
	return configs
}

func initConfig(filename string) (*viper.Viper, error) {

	v := viper.New()
	v.SetConfigFile(filename)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

func parseConfig(viperConfig *viper.Viper) (*Configurations, error) {
	var configs Configurations

	if err := viperConfig.Unmarshal(&configs); err != nil {
		return nil, err
	}
	return &configs, nil
}
