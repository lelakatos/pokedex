package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/lelakatos/pokedex/internal/pokecache"
)

func NewPokedex() map[string]PokeType {
	return map[string]PokeType{}
}

// TODO: finish function
func (c *Client) CatchPokemon(pokename string, csh *pokecache.Cache) (PokeType, error) {
	url := baseURL + "/pokemon/" + pokename
	if pokename == "" {
		return PokeType{}, errors.New("no poketype defined")
	}

	if pokemonData, exists := csh.Get(url); exists {
		var pokeData PokeType
		err := json.Unmarshal(pokemonData, &pokeData)
		if err != nil {
			return PokeType{}, err
		}

		return pokeData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeType{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeType{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeType{}, err
	}

	var pokeData PokeType
	err = json.Unmarshal(data, &pokeData)
	if err != nil {
		return PokeType{}, err
	}

	csh.Add(url, data)

	return pokeData, nil

}
