package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type main struct {
	Server   server   `yaml:"server"`
	Database database `yaml:"database"`
}

type server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var mainConfig *main

func Config() *main {
	return mainConfig
}

func DatbaseConfig() database {
	return mainConfig.Database
}

func ServerConfig() server {
	return mainConfig.Server
}

func LoadConfig(configFile string) (err error) {
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Loading config from file", viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(&mainConfig); err != nil {
		log.Fatalf("error unmarshilling config: %s\n", err.Error())
	}

	return
}
