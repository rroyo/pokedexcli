package main

import (
	"fmt"

	"github.com/rroyo/pokedexcli/internal/pokeapi"
)

func commandHelp(cfg *pokeapi.LocationArea) (*pokeapi.LocationArea, error) {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for cmd, command := range getCommands() {
		fmt.Printf("%s: %s\n", cmd, command.description)
	}

	fmt.Println("")

	return cfg, nil
}
