package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a valid pokemon name")
	}
	pokemonName := args[0]
	pokemonDetails, err := cfg.PokeApiClient.GetPokemonDetails(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)

	// Calculate the chance to catch pokemon
	max_base_experience := 635
	pokemon_base_experience := pokemonDetails.BaseExperience
	catch_chance := 100 - (pokemon_base_experience * 100 / max_base_experience)
	random_num := rand.Intn(101)
	if random_num < catch_chance {
		fmt.Printf("%v was caught!\n", pokemonName)
		// Add Pokemon into user's Pokedex
		cfg.caughtPokemon[pokemonName] = pokemonDetails
	} else {
		fmt.Printf("%v escaped!\n", pokemonName)
	}
	return nil
}
