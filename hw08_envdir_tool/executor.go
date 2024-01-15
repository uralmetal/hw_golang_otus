package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func RunCmd(cmd []string, env Environment) (returnCode int) {
	var err error
	var command *exec.Cmd

	for envName, envValue := range env {
		if !envValue.NeedRemove {
			err = os.Setenv(envName, envValue.Value)
		} else {
			err = os.Unsetenv(envName)
		}
		if err != nil {
			fmt.Println("Environment Error:", err)
			return 1
		}
	}
	if len(cmd) > 1 {
		command = exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	} else {
		command = exec.Command(cmd[0]) //nolint:gosec
	}
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin
	err = command.Run()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return exitError.ExitCode()
		}
		fmt.Println("Cmd run error:", err)
		return 1
	}
	return 0
}
