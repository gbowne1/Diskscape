package config

import (
	"fmt"
	"io/ioutil"
	// Keep this import
	"gopkg.in/yaml.v3"
)

type Config struct {
	TargetDir string `yaml:"target_dir"`
	Output    struct {
		File string `yaml:"file"`
	} `yaml:"output"`
}

var globalConfig *Config

func Load(path string) error {
	// Use ioutil.ReadFile here
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}

	globalConfig = &cfg
	return nil
}

func Get() *Config {
	if globalConfig == nil {
		panic("Configuration not loaded")
	}
	return globalConfig
}
