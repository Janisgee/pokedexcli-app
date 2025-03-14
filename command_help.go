package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	commands := getCommands()
	for i := range commands {
		fmt.Printf("%v: %v\n", commands[i].name, commands[i].description)
	}

	return nil
}
