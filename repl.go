package main

import (
	"strings"
)

type clicommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands map[string]clicommand
	next     *string
	prev     *string
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
	}
}

func cleanInput(text string) []string {
	text_lower := strings.ToLower(text)
	return strings.Fields(text_lower)
}
