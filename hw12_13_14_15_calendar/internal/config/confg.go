package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

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
