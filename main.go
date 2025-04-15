package main

func main() {
	myPokeDex := map[string]string{}
	commMap = map[string]cliCommand{
		"exit": {
			name:        "exit\t",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help\t",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map\t",
			description: "Shows the next 20 areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb\t",
			description: "Shows the previous 20 areas",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Show more information for the area you provide",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch\t",
			description: "Throw a pokeball at the pokemon you provide",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon you've already caught",
			callback:    commandInspect,
		},
	}
	PokeREPL(myPokeDex, commMap)
}
