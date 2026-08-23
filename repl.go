package main

import "strings"

func cleanInput(text string) []string {
	s := strings.ToLower(text)
	return strings.Fields(s)
}
