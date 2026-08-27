package main

import (
	"fmt"
)

func commandMapb(c *config, args []string) error {
	if c.prev == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url := *c.prev
	res, err := c.client.GetLocationAreas(url)
	if err != nil {
		return err
	}
	for _, v := range res.Results {
		fmt.Println(v.Name)
	}
	c.next = res.Next
	c.prev = res.Prev
	return nil
}
