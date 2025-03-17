package main

import "fmt"

func commandExplore(cfg *Config, args ...string) error {
	locationDetailRep, err := cfg.PokeApiClient.GetLocationPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v...\n", args[0])
	fmt.Println("Found Pokemon:")

	for _, v := range locationDetailRep.PokemonEncounters {
		fmt.Printf("- %v\n", v.Pokemon.Name)
	}

	return nil
}
