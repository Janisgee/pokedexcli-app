package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("there is no Pokemon have been caught")
	}
	for v := range cfg.caughtPokemon {
		fmt.Printf("- %v\n", v)
	}

	return nil
}
