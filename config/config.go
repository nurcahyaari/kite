package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Application struct {
		Name        string `mapstructure:"NAME"`
		Description string `mapstructure:"DESCRIPTION"`
		Version     string `mapstructure:"VERSION"`
	} `mapstructure:"APPLICATION"`
}

var cfg Config
var doOnce sync.Once

func Get() Config {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("cannot read config file")
	}

	doOnce.Do(func() {
		err := viper.Unmarshal(&cfg)
		if err != nil {
			log.Fatalln("cannot unmarshaling config")
		}
	})

	return cfg
}
