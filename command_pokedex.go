package main

import (
	"fmt"
)

func commandPokedex(cfg *config, _ []string) error {
	fmt.Print("Your Pokedex:\n")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf("  -%s\n", pokemon.Name)
	}

	return nil
}
