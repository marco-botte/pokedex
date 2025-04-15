package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokecache"
	"strings"
	"time"
)

func PokeREPL(myPokeDex map[string]Pokemon, commands map[string]cliCommand) {
	cache := pokecache.NewCache(time.Duration(20 * time.Second))
	next := "https://pokeapi.co/api/v2/location-area/"
	conf := config{
		Next:     &next,
		Previous: nil,
		Cache:    cache,
		Pokedex:  &myPokeDex,
	}
	var args []string = nil
	scanner := bufio.NewScanner(os.Stdin)
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
		command, ok := commands[input[0]]
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
