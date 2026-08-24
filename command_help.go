package main

import "fmt"

func commandHelp() error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
