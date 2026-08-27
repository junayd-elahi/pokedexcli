package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type PokemonEncounter struct {
	Pokemon Result `json:"pokemon"`
}
type LocationAreaDetail struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func (c *Client) GetLocationAreaDetail(url string) (LocationAreaDetail, error) {
	var data LocationAreaDetail

	body, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(body, &data)
		return data, err
	}
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaDetail{}, err
	}
	defer res.Body.Close()
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDetail{}, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return LocationAreaDetail{}, err
	}
	c.cache.Add(url, body)
	return data, nil
}
