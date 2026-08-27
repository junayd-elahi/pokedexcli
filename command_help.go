package main

import (
	"fmt"
)

func commandHelp(c *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, v := range c.commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}
