package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(c *config) error {
	if c.Previous != "" {
		res, err := http.Get(c.Previous)
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
	} else {
		fmt.Println("you're on the first page")
		return nil
	}
}
