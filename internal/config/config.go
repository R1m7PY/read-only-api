package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Timeout struct {
			Server time.Duration `yaml:"server"`
			Read   time.Duration `yaml:"read"`
			Idle   time.Duration `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"server"`
}

func NewConfig() (*Config, error) {
	config := Config{}

	file, err := os.Open("configs/app.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	yaml_decode := yaml.NewDecoder(file)

	if err := yaml_decode.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
