package main

import (
	"fmt"
)

func commandExplore(c *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("No area name Provided")
		return nil
	}
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]
	res, err := c.client.GetLocationAreaDetail(url)
	if err != nil {
		return err
	}

	fmt.Printf("\nExploring %v ...", args[0])
	fmt.Println("\n Found Pokemon:")
	for _, pokemon := range res.PokemonEncounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}
	return nil
}
