package config

import "encoding/json"

type Config struct {
	ServiceConfig *ServiceConfig
	StorageConfig *StorageConfig
}

type ServiceConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type StorageConfig struct {
	Conn string `yaml:"conn"`
}

func NewConfig() *Config {
	return &Config{
		ServiceConfig: &ServiceConfig{},
		StorageConfig: &StorageConfig{},
	}
}

func (c *Config) Dumps() (string, error) {
	d, err := json.MarshalIndent(c, "", "  ")
	return string(d), err
}
func (c *Config) Dump() ([]byte, error) {
	d, err := json.MarshalIndent(c, "", "  ")
	return d, err
}
