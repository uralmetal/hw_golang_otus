package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type LoggerConf struct {
	Level string `yaml:"level"`
	// TODO
}

func NewConfig(path string, config interface{}) error {
	err := ParseConfig(path, &config)
	return err
}

func ParseConfig(path string, config interface{}) error {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = parseYAML(yamlFile, config)
	if err != nil {
		return err
	}
	return nil
}

func parseYAML(content []byte, config interface{}) error {
	return yaml.Unmarshal(content, config)
}
