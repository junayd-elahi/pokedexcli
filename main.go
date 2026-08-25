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
		if len(textCleaned) > 0 {
			value, ok := cli()[textCleaned[0]]
			if ok {
				err := value.callback()
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
