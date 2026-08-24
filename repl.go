package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

		fmt.Printf("Your command was: %s\n", output[0])
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
