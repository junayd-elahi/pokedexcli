package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/junayde/pokedexcli/internal/pokecache"
)

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResponse struct {
	Next    *string  `json:"next"`
	Prev    *string  `json:"previous"`
	Results []Result `json:"results"`
}

type Client struct {
	cache *pokecache.Cache
}

func (c *Client) GetLocationAreas(url string) (LocationAreaResponse, error) {
	var data LocationAreaResponse

	body, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(body, &data)
		return data, err
	}
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer res.Body.Close()
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	c.cache.Add(url, body)
	return data, nil

}

func NewClient(interval time.Duration) *Client {
	return &Client{cache: pokecache.NewCache(interval)}
}
