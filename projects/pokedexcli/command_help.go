package main

import "fmt"

func callbackHelp(cfg *config, args...string) error {
	println("Welcome to the Pokedex help menu!")
	println("Available commands:")

	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Printf("  - %s: %s\n", cmd.name, cmd.description)
	}

	return nil
}
