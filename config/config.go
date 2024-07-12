package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DB struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Name     string `json:"name"`
	} `json:"db"`
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
	Queue struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Port     int    `json:"port"`
		Host     string `json:"host"`
	} `json:"queue"`
}

func LoadConfig(fileName string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
