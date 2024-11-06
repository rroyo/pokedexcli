package main

import (
	"os"

	"github.com/rroyo/pokedexcli/internal/pokeapi"
)

func commandExit(cfg *pokeapi.LocationArea) (*pokeapi.LocationArea, error) {
	os.Exit(0)
	return cfg, nil
}
