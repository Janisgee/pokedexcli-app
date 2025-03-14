package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Pokedex > ")
	reader := bufio.NewScanner(os.Stdin)

	// Create config
	cfg := &config{
		NextURL: "https://pokeapi.co/api/v2/location-area",
	}

	for reader.Scan() {
		input := reader.Text()
		trimInputList := cleanInput(input)

		if len(trimInputList) == 0 {
			continue
		}

		commandName := trimInputList[0]

		commands := getCommands()

		for i := range commands {
			if commandName == commands[i].name {
				err := commands[i].callback(cfg)
				if err != nil {

					fmt.Println(err)
					return
				}
			}
		}
		fmt.Print("Pokedex > ")
	}
}

func cleanInput(text string) []string {
	lowerCaseText := strings.ToLower(text)
	trimText := strings.TrimSpace(lowerCaseText)
	wordList := strings.Fields(trimText)
	return wordList
}
