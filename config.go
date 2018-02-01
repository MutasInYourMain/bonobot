package main

import (
	"os"
	"encoding/json"
)

// BonobotConfig contains all data that should be in config file
type BonobotConfig struct {
	Banwords []string `json:"banword_patterns"`
}

// ReadConfig reads config from config.json (or other file specified in envvar)
func ReadConfig() (BonobotConfig, error) {
	filename, exists := os.LookupEnv("BONOBOT_CONFIG")
	if !exists {
		filename = "config.json"
	}

	file, err := os.Open(filename)
	if err != nil {
		return BonobotConfig{}, err
	}

	decoder := json.NewDecoder(file)
	var config BonobotConfig
	err = decoder.Decode(&config)
	if err != nil {
		return BonobotConfig{}, err
	}
	return config, nil
}