package config

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ServiceConfig *ServiceConfig `yaml:"service_config"`
	StorageConfig *StorageConfig `yaml:"storage_config"`
}

type ServiceConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type StorageConfig struct {
	Conn string `yaml:"conn"`
}

func NewConfig() *Config {
	cfg := &Config{
		ServiceConfig: &ServiceConfig{},
		StorageConfig: &StorageConfig{},
	}
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func (c *Config) Dumps() (string, error) {
	d, err := json.MarshalIndent(c, "", "  ")
	return string(d), err
}
func (c *Config) Dump() ([]byte, error) {
	d, err := json.MarshalIndent(c, "", "  ")
	return d, err
}
