package main

import (
	"time"

	"github.com/ntino67/pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		registry:      getCommands(),
		pokeApiClient: pokeapi.NewClient(5*time.Second, 5*time.Second),
		pokedex:       make(map[string]pokeapi.Pokemon),
	}
	startREPL(cfg)
}
