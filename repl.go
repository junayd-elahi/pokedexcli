package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	text_lower := strings.ToLower(text)
	return strings.Fields(text_lower)
}

func commandExit() error {
	fmt.Errorf("Closing the Pokedex... Goodbye! ")
	os.Exit(0)
}
