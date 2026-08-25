package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ntino67/pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
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

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
		cache:      pokecache.NewCache(cacheInterval),
	}
}

func (c Client) ListLocation(url string) (RespShallowLocations, error) {
	jsonData, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, fmt.Errorf("network error: %w", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, fmt.Errorf("error getting from pokeapi: %w", err)
		}
		defer res.Body.Close()

		jsonData, err = io.ReadAll(res.Body)
		if err != nil {
			return RespShallowLocations{}, fmt.Errorf("error reading body: %w", err)
		}

		c.cache.Add(url, jsonData)
	}
	var locations RespShallowLocations

	if err := json.Unmarshal(jsonData, &locations); err != nil {
		return RespShallowLocations{}, fmt.Errorf("error unmarshalling json: %w", err)
	}

	return locations, nil
}
