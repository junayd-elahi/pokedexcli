package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Pick a pokemon to catch")
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	res, err := c.client.GetPokemon(args[0])
	if err != nil {
		return err
	}
	if rand.Intn(600) >= res.BaseExperience {
		fmt.Printf("%s was caught!", res.Name)
		c.pokedex[res.Name] = res
		return nil
	}
	fmt.Printf("%s escaped!", res.Name)
	return nil

}
