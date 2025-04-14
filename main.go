package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
