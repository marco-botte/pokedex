package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     *string
	Previous *string
}

var commMap map[string]cliCommand

func commandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range commMap {
		fmt.Println(command.name, "\t", command.description)
	}
	return nil
}

func commandExit(conf *config) error {
	defer os.Exit(0)
	fmt.Println("Closing the Pokedex... Goodbye!")
	return nil
}

func commandMap(conf *config) error {
	if conf.Next == nil {
		fmt.Println("No further areas! Go back.")
		return nil
	}
	locs, err := getLocs(*conf.Next)
	if err != nil {
		return err
	}
	conf.Next = locs.Next
	conf.Previous = locs.Previous
	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapBack(conf *config) error {
	if conf.Previous == nil {
		fmt.Println("No further areas! Go forward.")
		return nil
	}

	locs, err := getLocs(*conf.Previous)
	if err != nil {
		return err
	}
	conf.Next = locs.Next
	conf.Previous = locs.Previous
	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
