package main

import "fmt"

func commandHelp() error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for cmd, command := range getCommands() {
		fmt.Printf("%s: %s\n", cmd, command.description)
	}

	fmt.Println("")

	return nil
}
