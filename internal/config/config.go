package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

func Read() *Config {
	var config Config

	path, err := getConfigFilePath()
	if err != nil {
		reportError("Error finding path:", err)
	}

	file, err := os.Open(path + configFileName)
	if err != nil {
		reportError("Error opening file:", err)
	}

	defer file.Close()

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		reportError("Error parsing json:", err)
	}
	return &config
}

func (c *Config) SetUser(currentUser string) {
	c.CurrentUsername = currentUser
	if err := c.write(); err != nil {
		reportError("Error encoding to file:", err)
	}
}

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path, nil
}

func (c *Config) write() error {
	path, _ := getConfigFilePath()
	file, err := os.Create(path + configFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // pretty-print
	return encoder.Encode(c)
}

func reportError(msg string, err error) {
	fmt.Printf("%s: %v\n", msg, err)
	os.Exit(1)
}
