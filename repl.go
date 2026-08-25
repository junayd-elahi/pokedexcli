package main

import (
	"fmt"
	"os"
	"strings"
)

type clicommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands map[string]clicommand
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

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, v := range c.commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}
