package main

import (
	"fmt"
	"math/rand"
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
	Pokedex  *map[string]string
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

func commandCatch(conf *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Provide a pokemon to catch.")
		return nil
	}
	pokemon_name := args[0]
	pokemon, caught := (*conf.Pokedex)[pokemon_name]
	if caught {
		fmt.Printf("Already caught %s.\n", pokemon) // actual pokemon type useful?
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon_name)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemon_name)
	pokemonExp, err := getData[PokemonExperience](url, conf.Cache)
	if err != nil {
		return err
	}
	chance := 0.7 - min(float64(pokemonExp.Experience), 200)/300
	if chance > rand.Float64() {
		fmt.Printf("caught %s!\n", pokemon_name)
		(*conf.Pokedex)[pokemon_name] = pokemon_name // actual pokemon type useful?
	} else {
		fmt.Printf("%s got away.\n", pokemon_name)
	}
	return nil
}

func commandInspect(conf *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Provide a pokemon to inspect.")
		return nil
	}
	pokemon_name := args[0]
	poke, caught := (*conf.Pokedex)[pokemon_name]
	if !caught {
		fmt.Printf("%s not caught yet.\n", poke) // actual pokemon type useful?
		return nil
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemon_name)
	pokemon, err := getData[Pokemon](url, conf.Cache)
	if err != nil {
		return err
	}
	pokemon.Name = pokemon_name

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
