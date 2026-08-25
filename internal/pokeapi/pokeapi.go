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

type RespLocation struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel       int `json:"min_level"`
				PokemonDetails any `json:"pokemon_details"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func NewClient(timeout, cacheInterval time.Duration) *Client {
	return &Client{
		httpClient: http.Client{Timeout: timeout},
		cache:      pokecache.NewCache(cacheInterval),
	}
}

func (c *Client) ListLocation(url string) (RespShallowLocations, error) {
	jsonData, err := c.fetchData(url)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("error fetching data: %w", err)
	}

	var locations RespShallowLocations

	if err := json.Unmarshal(jsonData, &locations); err != nil {
		return RespShallowLocations{}, fmt.Errorf("error unmarshalling json: %w", err)
	}

	return locations, nil
}

func (c *Client) GetLocation(url string) (RespLocation, error) {
	jsonData, err := c.fetchData(url)
	if err != nil {
		return RespLocation{}, fmt.Errorf("error fetching data: %w", err)
	}

	var location RespLocation

	if err := json.Unmarshal(jsonData, &location); err != nil {
		return RespLocation{}, fmt.Errorf("error unmarshalling json: %w", err)
	}

	return location, nil
}

func (c *Client) fetchData(url string) ([]byte, error) {
	jsonData, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("network error: %w", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error getting from pokeapi: %w", err)
		}
		defer res.Body.Close()

		jsonData, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading body: %w", err)
		}

		c.cache.Add(url, jsonData)
	}

	return jsonData, nil
}
