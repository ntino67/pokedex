package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("no arguments provided")
	}
	for _, name := range args {
		url := "https://pokeapi.co/api/v2/pokemon/" + name

		pokemon, err := cfg.pokeApiClient.GetPokemon(url)
		if err != nil {
			return fmt.Errorf("error getting from pokeapi: %w", err)
		}

		fmt.Printf("Throwing a Pokeball at %s...\n", name)

		if rand.Intn(pokemon.BaseExperience) > 40 {
			fmt.Printf("%s escaped!\n", name)
			continue
		}

		fmt.Printf("%s was caught!\nYou can now inspect it with the inspect command", name)
		cfg.pokemon[name] = pokemon
	}

	return nil
}
