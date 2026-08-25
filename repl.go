package main

import (
	"fmt"
	"os"
	"strings"
)

type clicommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func cleanInput(text string) []string {
	text_lower := strings.ToLower(text)
	return strings.Fields(text_lower)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, v := range cli() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}
