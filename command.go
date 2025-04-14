package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commMap map[string]cliCommand

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range commMap {
		fmt.Println(command.name, ": ", command.description)
	}
	return nil
}

func commandExit() error {
	defer os.Exit(0)
	fmt.Println("Closing the Pokedex... Goodbye!")
	return nil
}
