package main

import (
	"fmt"

	"github.com/rroyo/pokedexcli/internal/pokeapi"
)

func commandMap(nav *pokeapi.LocationArea) (*pokeapi.LocationArea, error) {
	if nav.Next == nil {
		fmt.Println("End of areas")
		return nav, nil
	}

	loc, err := pokeapi.GetLocationAreas(nav.Next)
	if err != nil {
		return nav, err
	}

	for _, loc := range loc.Locations {
		fmt.Println(loc.Name)
	}

	return &loc, nil
}
