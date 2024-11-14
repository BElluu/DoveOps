package configuration

import (
	"encoding/json"
	"os"

	"github.com/rs/zerolog/log"
)

type Configuration struct {
	DatabaseConfig DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLmode  bool
	TimeZone string
}

func LoadConfiguration() (Configuration, error) {
	configFile := "config.json"
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
    log.Warn().Msg("Configuration file not found. Creating a new one.")
    newConfig := Configuration{
      DatabaseConfig: DatabaseConfig{
        Host:     "localhost",
        Port:     5432,
        User:     "postgres",
        Password: "password",
        Database: "dovedev",
        SSLmode:  false,
        TimeZone: "Warsaw/Poland",
      },
    }
    createConfigurationFile(newConfig)
	}
  file, err := os.Open(configFile)
  if err != nil {
    log.Error().Msg("Error opening configuration file.")
    return Configuration{}, err
  }

  defer file.Close()

  decoder := json.NewDecoder(file)
  var config Configuration
  err = decoder.Decode(&config)
  if err != nil {
    log.Error().Msg("Error decoding configuration file.")
    return Configuration{}, err
  }

  return config, nil
}

func createConfigurationFile(config Configuration) error {
  file, err := os.Create("config.json")
  if err != nil {
    log.Error().Msg("Error creating configuration file.")
    return err
  }

  defer file.Close()

  encoder := json.NewEncoder(file)
  encoder.SetIndent("", "  ")
  return encoder.Encode(config)
}
