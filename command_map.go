package main

import (
	"errors"
	"fmt"
)

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Locations struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

func commandMap(cfg *config) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
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

func commandMapb(cfg *config) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("You are on the first page")
	}
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
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
