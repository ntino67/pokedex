package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ntino67/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	registry      map[string]cliCommand
	next          *string
	previous      *string
	pokeApiClient pokeapi.Client
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {name: "exit", description: "Exit the pokedex", callback: commandExit},
		"help": {name: "help", description: "Displays a help message", callback: commandHelp},
		"map":  {name: "map", description: "Displays the next 20 locations", callback: commandMap},
		"mapb": {name: "mapb", description: "Displays the previous 20 locations", callback: commandMapBack},
	}
}

func cleanInput(text string) []string {
	s := strings.ToLower(text)
	return strings.Fields(s)
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		output := cleanInput(scanner.Text())

		if len(output) == 0 {
			continue
		}

		command := output[0]

		cmd, exists := cfg.registry[command]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.callback(cfg); err != nil {
			fmt.Println("Error:", err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
