package main

import (
	"bufio"
	"fmt"

	// "io"
	"os"
)

func main() {
	for {
		fmt.Print("Pokedex >")
		scanner := bufio.NewScanner(os.Stdin)
		ok := scanner.Scan()
		if !ok {
			fmt.Println("Error reading string")
		}

		inputText := scanner.Text()

		words := cleanInput(inputText)
		if len(words) == 0 {
			fmt.Println("You didn't input anything, cheeky")
		}

		fmt.Printf("Your command was: %s\n", words[0])

	}
}
