package main

import (
	"fmt"

	"github.com/janisgee/pokedexcli-app/internal/pokeapi"
)

type ResourceLocationName struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *config) error {
	locationRep, err := pokeapi.GetLocation(cfg.NextURL)
	if err != nil {
		return fmt.Errorf("error getting location name list.")
	}

	// set config pagination of next
	cfg.PreviousURL = locationRep.Previous
	cfg.NextURL = locationRep.Next

	for _, v := range locationRep.Results {
		fmt.Printf("%v\n", v.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	locationRep, err := pokeapi.GetLocation(cfg.PreviousURL)
	if err != nil {
		return fmt.Errorf("error getting location name list.")
	}

	// set config pagination of next
	cfg.PreviousURL = locationRep.Previous
	cfg.NextURL = locationRep.Next

	for _, v := range locationRep.Results {
		fmt.Printf("%v\n", v.Name)
	}

	return nil
}
