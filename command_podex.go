package main

import (
	"fmt"
)

func commandPokedex(c *config, args []string) error {
	if len(c.pokedex) == 0 {
		fmt.Println("No Pokemon in Pokedex")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, v := range c.pokedex {
		fmt.Printf(" - %s\n", v.Name)
	}
	return nil
}
