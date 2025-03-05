package main

import (
	"fmt"
	"math/rand/v2"
)

// TODO finish function
func commandCatch(cfg *config, pokeName string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokeName)
	pokeData, err := cfg.pokeapiClient.CatchPokemon(pokeName, cfg.cache)
	if err != nil {
		return err
	}

	if rand.Float64()*float64(pokeData.BaseExperience) < 50 {
		fmt.Printf("Gotcha! Successfully caught %s.\n", pokeName)
		cfg.pokedex[pokeName] = pokeData
	} else {
		fmt.Printf("Better luck next time. %s escaped.\n", pokeName)
	}

	return nil
}
