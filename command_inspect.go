package main

import "fmt"

func commandInspect(cfg *config, pokename string) error {
	pokemon, exists := cfg.pokedex[pokename]
	if !exists {
		fmt.Printf("You haven't caught %s yet\n", pokename)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  - %s: %v\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
