package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string `json:"address"`
}

type Config struct {
	Env         string `json:"env"`
	StoragePath string `json:"storagepath"`
	HttpServer  `json:"httpserver"`
}

func MustLoad() *Config {

	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {

		flags := flag.String("env", "", "environment variable for configpath")
		flag.Parse()
		configPath = *flags
		if configPath == "" {
			log.Fatal("Please set the config path and then continue")
		}

	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s provided not found", configPath)

	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Structure in the config file is not supported: %v", err.Error())
	}
	return &cfg
}
