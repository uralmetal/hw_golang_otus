package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()
	if from == "" {
		fmt.Println("Parameter 'from' is required")
		os.Exit(1)
	}
	if to == "" {
		fmt.Println("Parameter 'to' is required")
		os.Exit(1)
	}
	fmt.Println("Run copy file from ", from, " to ", to, " with limit=", limit, " and offset=", offset)
	err := Copy(from, to, limit, offset)
	if err == nil {
		fmt.Println("Runtime error:\n", err)
		os.Exit(1)
	}
	fmt.Println("File ", from, " successfully copied to ", to)
}
