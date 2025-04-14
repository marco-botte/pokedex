package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokecache"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	myPokeDex := map[string]string{}
	commMap = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Shows the next 20 areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 areas",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Show more information for the area you provide",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Throw a pokeball at the pokemon you provide",
			callback:    commandCatch,
		},
	}

	cache := pokecache.NewCache(time.Duration(20 * time.Second))
	next := "https://pokeapi.co/api/v2/location-area/"
	conf := config{
		Next:     &next,
		Previous: nil,
		Cache:    cache,
		Pokedex:  &myPokeDex,
	}
	var args []string = nil
	for {
		args = []string{}
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		input := cleanInput(line)
		if len(input) == 0 {
			continue
		}
		if len(input) > 1 {
			args = input[1:2]
		}
		command, ok := commMap[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.callback(&conf, args...)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}

}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words
}
