package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	var pokemon Pokemon

	body, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(body, &pokemon)
		return pokemon, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, body)
	return pokemon, nil
}
