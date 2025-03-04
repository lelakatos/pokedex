package main

import (
	"time"

	"github.com/lelakatos/pokedex/internal/pokeapi"
	"github.com/lelakatos/pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	csh := pokecache.NewCache(30 * time.Second)

	cfg := &config{
		pokeapiClient: pokeClient,
		cache:         csh,
	}

	startRepl(cfg)
}
