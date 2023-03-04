package config

import (
	"github.com/spf13/viper"
	"log"
)

type Configuration struct {
	ModulePath string `json:"module_path"`
}

var (
	Config *Configuration
)

func init() {
	viper.SetDefault("ModulePath", "modules")

	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("WARN - No config file found, using defaults")
		return
	}

	err = viper.Unmarshal(Config)
	if err != nil {
		panic(err)
	}
}
