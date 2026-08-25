package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cfg := config{
		commands: cli(),
	}
	startRepl(&cfg)
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\nPokedex > ")
		scanner.Scan()
		text := scanner.Text()
		textCleaned := cleanInput(text)
		if len(textCleaned) > 0 {
			value, ok := cfg.commands[textCleaned[0]]
			if ok {
				err := value.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
