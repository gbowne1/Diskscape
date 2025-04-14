package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Config represents the structure of the YAML configuration file
type Config struct {
	TargetDir string `yaml:"target_dir"`
	Output    struct {
		File string `yaml:"file"`
	} `yaml:"output"`
	Report struct {
		MaxEntries int `yaml:"max_entries"`
	} `yaml:"report"`
	Thresholds struct {
		WarningPercentage  int `yaml:"warning_percentage"`
		CriticalPercentage int `yaml:"critical_percentage"`
	} `yaml:"thresholds"`
	Logging struct {
		Level string `yaml:"level"`
		File  string `yaml:"file"`
	} `yaml:"logging"`
}

// GlobalConfig holds the parsed configuration values
var GlobalConfig *Config

// Load reads and parses the YAML configuration file
func Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}

	GlobalConfig = &cfg
	log.Printf("Configuration loaded successfully from %s\n", path)
	return nil
}

// Get returns the loaded configuration
func Get() *Config {
	if GlobalConfig == nil {
		panic("Configuration not loaded")
	}
	return GlobalConfig
}
