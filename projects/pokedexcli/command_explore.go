package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}

	LocationAreaName := args[0]

	LocationArea, err := cfg.pokeapiClient.GetLocationArea(
		LocationAreaName,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s:\n", LocationArea.Name)
	for _, pokemon := range LocationArea.PokemonEncounters {
		fmt.Printf("  - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
