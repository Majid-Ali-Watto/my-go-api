package configs

import (
	"log"
	"os"

	"io"

	"gopkg.in/yaml.v2"
)

// Config holds the application configuration
type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

var AppConfig Config

// LoadConfig loads the configuration from a YAML file
func LoadConfig() {
	// Open the config file
	file, err := os.Open("configs/config.yaml")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	// Read the file contents
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Unmarshal the YAML data
	if err := yaml.Unmarshal(data, &AppConfig); err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	// Check for environment variable
	if port := os.Getenv("PORT"); port != "" {
		AppConfig.Server.Port = port
	}
}
