package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/lelakatos/pokedex/internal/pokecache"
)

//List Locations -

func (c *Client) ListLocations(pageURL *string, csh *pokecache.Cache) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cachedValue, exists := csh.Get(url); exists {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(cachedValue, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	csh.Add(url, dat)

	return locationsResp, nil
}
