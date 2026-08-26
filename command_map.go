package main

import (
	"fmt"

	"github.com/junayde/pokedexcli/internal/pokeapi"
)

func commandMap(c *config) error {
	var url string
	if c.next == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = *c.next
	}
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
