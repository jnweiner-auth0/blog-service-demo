package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

var Conf *Configuration

type Configuration struct {
	Port     int    `envconfig:"PORT" default:"5050"`
	DBEngine string `envconfig:"DB_ENGINE" default:"mongo"`
}

func GetConfiguration() *Configuration {
	return Conf
}

func init() {
	var c Configuration
	err := envconfig.Process("", &c)
	if err != nil {
		fmt.Printf("err in config, %v", err)
		panic("error initializing config")
	}
	Conf = &c
}
