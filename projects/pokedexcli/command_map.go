package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(
		cfg.nextLocationAreaURL,
	)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("  - %s: %s\n", area.Name, area.URL)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	return nil
}

func callbackMapB(cfg *config, args...string) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you are on the first page")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(
		cfg.prevLocationAreaURL,
	)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("  - %s: %s\n", area.Name, area.URL)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	return nil
}
