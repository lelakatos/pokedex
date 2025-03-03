package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl() {
	for {
		fmt.Print("Pokedex >")
		scanner := bufio.NewScanner(os.Stdin)
		ok := scanner.Scan()
		if !ok {
			fmt.Println("Error reading string")
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if v, ok := commandList[commandName]; ok {
			err := v.callback()
			if err != nil {
				fmt.Println("error:", err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}

}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandList map[string]cliCommand = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}
