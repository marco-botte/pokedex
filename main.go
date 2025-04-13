package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		command.callback()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}
