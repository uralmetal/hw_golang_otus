package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Number of arguments too small")
	}
	env, err := ReadDir(args[1])
	if err != nil {
		fmt.Println("Read env error:", err)
	}
	returnCode := RunCmd(args[2:], env)
	os.Exit(returnCode)
}
