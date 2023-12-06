package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/example/hello/reverse"
)

func SayHelloOtus(w io.Writer) {
	const helloPhrase = "Hello, OTUS!"

	fmt.Fprintf(w, "%s\n", reverse.String(helloPhrase))
}

func main() {
	SayHelloOtus(os.Stdout)
}
