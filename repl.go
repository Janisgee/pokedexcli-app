package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/janisgee/pokedexcli-app/internal/pokeapi"
)

type Config struct {
	PokeApiClient pokeapi.Client
	caughtPokemon map[string]pokeapi.Pokemon
	NextURL       string
	PreviousURL   string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func startPokedex(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := reader.Text()
		trimInputList := cleanInput(input)

		if len(trimInputList) == 0 {
			//Invalid input
			continue
		}

		commandName := trimInputList[0]
		args := []string{}
		if len(trimInputList) > 1 {
			args = trimInputList[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
				continue
			}
			continue
		} else {
			fmt.Println("Unknown Command")
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"pokedex": {
			name:        "pokedex",
			description: "Track all the Pokemon user have caught",
			callback:    commandPokedex,
		},

		"explore": {
			name:        "explore",
			description: "Displays a list of Pokemon located in the input location",
			callback:    commandExplore,
		},

		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},

		"inspect": {
			name:        "inspect",
			description: "inspect a pokemon",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},

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
	}
}
