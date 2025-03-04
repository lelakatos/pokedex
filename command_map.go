package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, p string) error {

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL, cfg.cache)
	if err != nil {
		return fmt.Errorf("error with the API call: %w", err)
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.previousLocationsURL = locationResp.Previous

	for _, result := range locationResp.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(cfg *config, p string) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("you are on the first page")
	}
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL, cfg.cache)
	if err != nil {
		return fmt.Errorf("error with the API call: %w", err)
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.previousLocationsURL = locationResp.Previous

	for _, result := range locationResp.Results {
		fmt.Println(result.Name)
	}

	return nil
}
