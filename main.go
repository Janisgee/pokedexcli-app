package main

import (
	"strings"
	"time"

	"github.com/janisgee/pokedexcli-app/internal/pokeapi"
)

func main() {

	// Create Config - httpClient and cache
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &Config{
		PokeApiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startPokedex(cfg)

}

func cleanInput(text string) []string {
	lowerCaseText := strings.ToLower(text)
	trimText := strings.TrimSpace(lowerCaseText)
	wordList := strings.Fields(trimText)
	return wordList
}
