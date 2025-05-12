package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("(pokedex) > ")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Printf("Unknown command: %s\n", commandName)
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help | h",
			description: "Shows the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit | quit",
			description: "Turns off the Pokedex CLI",
			callback:    callbackExit,
		},
		"map": {
			name:        "map | m",
			description: "List the next page of location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb | mb",
			description: "List the previous page of location areas",
			callback:    callbackMapB,
		},
		"explore": {
			name:        "explore {location area}",
			description: "Explore a location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon}",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon}",
			description: "View information a pokemon that has been added to your pokedex",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex | p",
			description: "List the pokemon in your pokedex",
			callback:    callbackPokedex,
		},
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)

	words := strings.Fields(lowered)

	return words
}
