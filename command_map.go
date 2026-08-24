package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.next != nil {
		url = *cfg.next
	}

	locations, err := cfg.pokeApiClient.ListLocation(url)
	if err != nil {
		return fmt.Errorf("error getting from pokeapi: %w", err)
	}

	cfg.next = locations.Next
	cfg.previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}
