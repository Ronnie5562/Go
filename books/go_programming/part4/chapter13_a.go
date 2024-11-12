package main

import (
	"flag"
	"fmt"
	// "os"
)

func main() {
	fmt.Println("Programming from the command line")

	// args := os.Args
	// if len(args) < 2 {
	// 	fmt.Println("Please provide an argument")
	// 	return
	// }
	// name := args[1]
	// fmt.Printf("Hello, %s! \nWelcome to Ronnie's CLI\n", name)


	UseFlags()
}

// To see the list of flags that can be used, run the command below:
// go run chapter13.go --help

var (
	nameFlag = flag.String("name", "Maya", "The name of the user")
	quietFlag = flag.Bool("quiet", false, "Toggle to greet the user quietly")
)

func UseFlags() {
	flag.Parse()

	if !*quietFlag {
		fmt.Printf("Hello, %s! \nWelcome to Ronnie's CLI", *nameFlag)
	} else {
		fmt.Printf("Welcome to Ronnie's CLI")
	}
}
