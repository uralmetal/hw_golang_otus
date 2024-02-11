package main

import (
	configHandler "github.com/uralmetal/hw_golang_otus/hw12_13_14_15_calendar/internal/config"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type LoggerConf struct {
	Level string `yaml:"level"`
	// TODO
}

type Config struct {
	Logger LoggerConf `yaml:"logger"`
	// TODO
}

func NewConfig(path string) (Config, error) {
	var config Config
	err := configHandler.ParseConfig(path, &config)
	return config, err
}

// TODO
