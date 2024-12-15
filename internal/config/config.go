package config

import "os"

type Config struct {
	HTTPAddr string
}

func Read() Config {
	var config Config
	httpAddr, exists := os.LookupEnv("HTTP_ADDR")
	if exists {
		config.HTTPAddr = httpAddr
	}
	return config
}
