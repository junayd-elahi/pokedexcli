package main

import (
	"fmt"

	"github.com/junayde/pokedexcli/internal/pokeapi"
)

func commandMapb(c *config) error {
	if c.prev == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url := *c.prev
	res, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}
	for _, v := range res.Results {
		fmt.Println(v.Name)
	}
	c.next = res.Next
	c.prev = res.Prev
	return nil
}
