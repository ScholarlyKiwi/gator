package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBurl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	var gatorconfig Config

	configfile, err := getConfigFilePath()
	if err != nil {
		return gatorconfig, fmt.Errorf("Read Config %v", err)
	}

	fileRead, err := os.ReadFile(configfile)
	if err != nil {
		return gatorconfig, fmt.Errorf("Read Config Error reading %v: %v", configfile, err)
	}

	if err := json.Unmarshal(fileRead, &gatorconfig); err != nil {
		return gatorconfig, fmt.Errorf("Read Config Error reading config json: %v", err)
	}
	return gatorconfig, nil

}

func (gatorConfig *Config) SetUser(username string) error {
	gatorConfig.CurrentUserName = username
	write(*gatorConfig)
	return nil
}

func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("Write Config %v", err)
	}

	byteArray, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Write Config Error marshalling Json: %v", err)
	}

	err = os.WriteFile(configFilePath, byteArray, 0644)
	if err != nil {
		return fmt.Errorf("Write Config Error writing file: %v", err)
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error reading Config: %v", err)
	}

	configfile := path.Join(homeDir, configFileName)
	if err != nil {
		return "", fmt.Errorf("Error finding config file name: %v", err)
	}

	return configfile, nil
}
