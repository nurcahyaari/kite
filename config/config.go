package config

import (
	"sync"
)

type Config struct {
	Application struct {
		Name        string
		Description string
		Version     string
	}
}

var cfg Config
var doOnce sync.Once

func Get() Config {
	doOnce.Do(func() {
		cfg = Config{
			Application: struct {
				Name        string
				Description string
				Version     string
			}{
				Name:        "Kite",
				Description: "Kite is a golang code structure generator. It is easy to use and helps to boost your productivity. no need to recreate your code structure when creating new services, or apps",
				Version:     "0.5.1",
			},
		}
	})

	return cfg
}
