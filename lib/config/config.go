package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "config/local.yaml", "Path to the configuration file")
		flag.Parse()

		configPath = *flags
	}

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set and no default config file provided")
	}

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Fatalf("Configuration file does not exist: %s", configPath)
	}

	var config Config

	nerr := cleanenv.ReadConfig(configPath, &config)

	if nerr != nil {
		log.Fatalf("Failed to read configuration file: %s, error: %v", configPath, nerr.Error())
	}

	return &config
}
