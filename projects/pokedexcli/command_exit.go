package main

import "os"

func callbackExit(cfg *config, args...string) error {
	println("Exiting the Pokedex CLI...")
	os.Exit(0)

	return nil
}
