package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	locationRep, err := cfg.PokeApiClient.ListLocation(cfg.NextURL)
	if err != nil {
		return fmt.Errorf("error getting location name list from fetching pokeapi")
	}

	// set config pagination of next
	cfg.PreviousURL = locationRep.Previous
	cfg.NextURL = locationRep.Next

	for _, v := range locationRep.Results {
		fmt.Printf("%v\n", v.Name)
	}

	return nil
}

func commandMapb(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	locationRep, err := cfg.PokeApiClient.ListLocation(cfg.PreviousURL)
	if err != nil {
		return fmt.Errorf("error getting location name list")
	}

	// set config pagination of next
	cfg.PreviousURL = locationRep.Previous
	cfg.NextURL = locationRep.Next

	for _, v := range locationRep.Results {
		fmt.Printf("%v\n", v.Name)
	}

	return nil
}
