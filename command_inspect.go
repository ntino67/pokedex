package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("no arguments provided")
	}
	for _, name := range args {
		pokemon, ok := cfg.pokedex[name]
		if !ok {
			fmt.Printf("The pokemon %s is not in the pokedex...\nTry to catch it!\n", name)
		}

		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Print("Stats:\n")
		for _, s := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
		}
		fmt.Print("Types:\n")
		for _, t := range pokemon.Types {
			fmt.Printf("  -%s\n", t.Type.Name)
		}
	}

	return nil
}
