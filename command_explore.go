package main

import "fmt"

// TODO: Finish function
func commandExplore(cfg *config, area string) error {
	fmt.Printf("Exploring %s...\n", area)
	exploreOptions, err := cfg.pokeapiClient.ExploreLocations(area, cfg.cache)
	if err != nil {
		return fmt.Errorf("error with the API call: %w", err)
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range exploreOptions.PokemonEncounters {
		fmt.Println("- ", pokemon.Pokemon.Name)
	}

	return nil
}
