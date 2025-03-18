package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name for inspection")
	}

	pokemonName := args[0]
	// Check if you have caught pokemon
	details, exist := cfg.caughtPokemon[pokemonName]
	if exist {
		fmt.Printf("Name: %v\n", pokemonName)
		fmt.Printf("Height: %v\n", details.Height)
		fmt.Printf("Weight: %v\n", details.Weight)
		fmt.Println("Stats:")
		// Loop Stats and display all the stats values
		for _, val := range details.Stats {
			fmt.Printf("  - %v: %v\n", val.Stat.Name, val.BaseStat)
		}
		fmt.Println("Types:")
		// Loop Types and display all the types values
		for _, val := range details.Types {
			fmt.Printf("  - %v\n", val.Type.Name)
		}
		return nil
	} else {
		return errors.New("you have not caught that pokemon")
	}
}
