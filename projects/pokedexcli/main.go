package main

import (
	"time"

	"github.com/ronnie5562/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient:       pokeapi.NewClient(time.Hour),
		caughtPokemon:       make(map[string]pokeapi.Pokemon),
		nextLocationAreaURL: nil,
		prevLocationAreaURL: nil,
	}

	startRepl(&cfg)
}
