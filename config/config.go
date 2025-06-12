package config

import (
	"log"
	"os"

	"github.com/Ik-cyber/caching-proxy/models"
	"gopkg.in/yaml.v3"
)

func LoadConfig(filename string) *models.Config {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
	var cfg models.Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}
	return &cfg
}
