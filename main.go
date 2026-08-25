package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\nPokedex > ")
		scanner.Scan()
		text := scanner.Text()
		textCleaned := cleanInput(text)
		fmt.Printf("Your command was: %v", textCleaned[0])
	}
}
