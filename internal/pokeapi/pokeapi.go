package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
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

func GetLocationAreas(url string) (LocationAreaResponse, error) {
	var data LocationAreaResponse
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	return data, nil
}
