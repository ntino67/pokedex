package main

import (
	"pokeapi"
	"time"
)

func main() {
	cfg := &config{
		registry: getCommands(),
		pokeApiClient: pokeapi.NewClient(5 * time.Second),
	}
	startREPL(cfg)
}
