package main

import (
	"strings"

	"github.com/junayde/pokedexcli/internal/pokeapi"
)

type clicommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	commands map[string]clicommand
	next     *string
	prev     *string
	client   *pokeapi.Client
	pokedex  map[string]pokeapi.Pokemon
}

func cli() map[string]clicommand {
	return map[string]clicommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Shows the map of the area",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Goes back one page for the map",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores the map",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catches the Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspects the Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows the pokemon",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	text_lower := strings.ToLower(text)
	return strings.Fields(text_lower)
}
