package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/junayde/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := config{
		commands: cli(),
		client:   pokeapi.NewClient(5 * time.Second),
		pokedex:  map[string]pokeapi.Pokemon{},
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
				err := value.callback(cfg, textCleaned[1:])
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
