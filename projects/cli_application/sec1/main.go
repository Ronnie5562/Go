package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	greeting(os.Stdout, "John")
}

func greeting(out io.Writer, name string) {
	fmt.Fprintf(out, "Hello, %s!\n", name)
}
