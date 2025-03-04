package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/lelakatos/pokedex/internal/pokecache"
)

// TODO add cache functionality
func (c *Client) ExploreLocations(location string, csh *pokecache.Cache) (RespDetailedLocations, error) {
	url := baseURL + "/location-area/" + location
	if location == "" {
		return RespDetailedLocations{}, errors.New("Location does not exist")
	}

	if cachedValue, exists := csh.Get(url); exists {
		exploreResp := RespDetailedLocations{}
		err := json.Unmarshal(cachedValue, &exploreResp)
		if err != nil {
			return RespDetailedLocations{}, err
		}
		return exploreResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDetailedLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDetailedLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDetailedLocations{}, nil
	}

	var exploreResp RespDetailedLocations
	err = json.Unmarshal(data, &exploreResp)
	if err != nil {
		return RespDetailedLocations{}, err
	}

	csh.Add(url, data)

	return exploreResp, nil
}
