package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Stat struct {
	BaseStat int    `json:"base_stat"`
	Stat     Result `json:"stat"`
}

type Type struct {
	Slot int    `json:"slot"`
	Type Result `json:"type"`
}

type Pokemon struct {
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`
	Stats          []Stat `json:"stats"`
	Types          []Type `json:"types"`
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
