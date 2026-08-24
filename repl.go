package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {name: "exit", description: "Exit the pokedex", callback: commandExit,},
		"help": {name: "help", description: "Displays a help message", callback: commandHelp,},
	}
}

func cleanInput(text string) []string {
	s := strings.ToLower(text)
	return strings.Fields(s)
}

func startREPL() {
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
		commands := getCommands()

		if commands[command].name == "" {
			fmt.Println("Unknown command")
			continue
		}

		if err := commands[command].callback(); err != nil {
			fmt.Println("Error:", err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
