package main

import (
	"fmt"
	"os"
)

func commandExit(c *config, p string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
