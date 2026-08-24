package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
	}
}

func (c Client) ListLocation(url string) (RespShallowLocations, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("network error: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("error getting from pokeapi: %w", err)
	}
	defer res.Body.Close()
	
	var locations RespShallowLocations
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return RespShallowLocations{}, fmt.Errorf("error decoding json: %w", err)
	}

	return locations, nil
}
