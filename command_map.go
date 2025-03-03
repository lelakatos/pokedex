package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func commandMap(c *config) error {
	res, err := http.Get(c.Next)
	if err != nil {
		return fmt.Errorf("error with the API call: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status call: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("reading response body failed: %w", err)
	}

	locations := Locations{}

	err = json.Unmarshal(body, &locations)
	if err != nil {
		return err
	}

	c.Next = locations.Next
	c.Previous = locations.Previous

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return nil
}
