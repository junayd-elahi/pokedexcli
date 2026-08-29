package main

import (
	"fmt"
)

func commandInspect(c *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Enter a Pokemon name to inspect")
		return nil
	}
	value, ok := c.pokedex[args[0]]
	if !ok {
		fmt.Println("you have not caught that Pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", value.Name)
	fmt.Printf("Height: %d\n", value.Height)
	fmt.Printf("Weight: %d\n", value.Weight)
	fmt.Println("Stats:")
	for _, v := range value.Stats {
		fmt.Printf("  -%s: %d\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range value.Types {
		fmt.Printf("  - %s\n", v.Type.Name)
	}
	return nil
}
