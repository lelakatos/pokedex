package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lelakatos/pokedex/internal/pokeapi"
	"github.com/lelakatos/pokedex/internal/pokecache"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		var paramName string

		if len(words) > 1 {
			paramName = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, paramName)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List next page of location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List prev page of location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Returns all the pokemon located in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch the named pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon that has already been caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See all the pokemon you caught",
			callback:    commandPokedex,
		},
	}
}

type config struct {
	pokeapiClient        pokeapi.Client
	cache                *pokecache.Cache
	pokedex              map[string]pokeapi.PokeType
	nextLocationsURL     *string
	previousLocationsURL *string
}
