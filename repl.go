package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rroyo/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.LocationArea) (*pokeapi.LocationArea, error)
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	nextURL := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	var previousURL *string

	cfg := pokeapi.LocationArea{
		Next:      &nextURL,
		Previous:  previousURL,
		Locations: nil,
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			nav, err := command.callback(&cfg)
			if err != nil {
				fmt.Println(err)
				continue
			}

			cfg.Next = nav.Next
			cfg.Previous = nav.Previous

		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show the next 20 areas",
			callback:    commandMap,
		},
		"bmap": {
			name:        "bmap",
			description: "Show the previous 20 areas",
			callback:    commandBmap,
		},
	}
}
