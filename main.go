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
	}
	cache := pokecache.NewCache(time.Duration(20 * time.Second))
	next := "https://pokeapi.co/api/v2/location-area/"
	conf := config{
		Next:     &next,
		Previous: nil,
		Cache:    cache,
	}
	for {

		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		input := cleanInput(line)
		if len(input) == 0 {
			continue
		}
		command, ok := commMap[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.callback(&conf)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}

}

func cleanInput(text string) []string {
	return strings.Fields(text)
}
