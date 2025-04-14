package main

import (
	"fmt"
	"os"
	"pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	Next     *string
	Previous *string
	Cache    *pokecache.Cache
}

var commMap map[string]cliCommand

func commandHelp(conf *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range commMap {
		fmt.Println(command.name, "\t", command.description)
	}
	return nil
}

func commandExit(conf *config, args ...string) error {
	defer os.Exit(0)
	fmt.Println("Closing the Pokedex... Goodbye!")
	return nil
}

func commandMap(conf *config, args ...string) error {
	if conf.Next == nil {
		fmt.Println("No further areas! Go back.")
		return nil
	}
	locs, err := getData[LocAreaResponse](*conf.Next, conf.Cache)
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

func commandMapBack(conf *config, args ...string) error {
	if conf.Previous == nil {
		fmt.Println("No further areas! Go forward.")
		return nil
	}

	locs, err := getData[LocAreaResponse](*conf.Previous, conf.Cache)
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

func commandExplore(conf *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Provide an area to explore.")
		return nil
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", args[0])
	location, err := getData[LocationResponse](url, conf.Cache)
	if err != nil {
		return err
	}
	for _, encounter := range location.Encounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
