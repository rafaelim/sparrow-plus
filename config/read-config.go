package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port string `json:"port"`
	RootDir string `json:"rootDir"`
}

func ReadConfig() *Config {
	encondedConfig, err := os.ReadFile("./config.json")

	if err != nil {
		log.Fatalf("No config file was found %v", err)
	}
	var config *Config
	converr := json.Unmarshal(encondedConfig, &config)

	if converr != nil {
		log.Fatalf("Failed to convert config file %v", converr)
	}

	return config
}