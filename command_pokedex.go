package main

import "fmt"

func commandPokedex(cfg *config, s string) error {
	fmt.Println("Your pokedex:")

	if len(cfg.pokedex) == 0 {
		fmt.Println("Your pokedex is empty!")
		return nil
	}

	for poke, _ := range cfg.pokedex {
		fmt.Printf("  - %s\n", poke)
	}

	return nil
}
