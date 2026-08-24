package main

import "fmt"

func commandMapBack(cfg *config) error {
	if cfg.previous == nil {
		return fmt.Errorf("no previous locations")
	}
	url := *cfg.previous

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

