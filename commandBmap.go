package main

import (
	"fmt"

	"github.com/rroyo/pokedexcli/internal/pokeapi"
)

func commandBmap(nav *pokeapi.LocationArea) (*pokeapi.LocationArea, error) {
	if nav.Previous == nil {
		fmt.Println("Can't go further back")
		return nav, nil
	}

	loc, err := pokeapi.GetLocationAreas(nav.Previous)
	if err != nil {
		return nav, err
	}

	for _, loc := range loc.Locations {
		fmt.Println(loc.Name)
	}

	return &loc, nil
}
