package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("no arguments provided")
	}
	for _, name := range args {
		url := "https://pokeapi.co/api/v2/location-area/" + name

		location, err := cfg.pokeApiClient.GetLocation(url)
		if err != nil {
			return fmt.Errorf("error getting from pokeapi: %w", err)
		}

		fmt.Printf("Exploring %s...\nFound Pokemon:\n", name)

		for _, pokemon := range location.PokemonEncounters {
			fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
		}
	}

	return nil
}
