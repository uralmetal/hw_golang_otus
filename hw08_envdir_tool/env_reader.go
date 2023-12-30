package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

func readFile(path string) (EnvValue, error) {
	valueRaw, err := os.ReadFile(path)
	if err != nil {
		return EnvValue{}, err
	}
	needRemove := len(valueRaw) == 0
	valueLines := strings.SplitN(string(valueRaw), "\n", 2)
	valueStr := strings.ReplaceAll(valueLines[0], "\x00", "\n")
	valueStr = strings.TrimRight(valueStr, " \t")

	return EnvValue{
		Value:      valueStr,
		NeedRemove: needRemove,
	}, nil
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	env := make(Environment)
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileName := file.Name()
		if strings.ContainsRune(fileName, '=') {
			fmt.Println("Skip invalid filename", fileName, "in dir", dir)
			continue
		}
		value, err := readFile(filepath.Join(dir, fileName))
		if err != nil {
			fmt.Println("Error read file", fileName, ":", err)
			continue
		}
		env[fileName] = value
	}
	return env, nil
}
