package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRelp(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()

		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		commands := getCommands()

		userCommand := words[0]
		command, ok := commands[userCommand]

		if !ok {
			fmt.Println("Invalid command.")
			continue
		}

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		err := command.callback(conf, args...)

		if err != nil {
			fmt.Println(err)
		}

	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations areas in the Pokemon World",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations ares in the Pokemon World",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore {location}",
			description: "Give the list of Pokemon that a appear on a given area",
			callback:    commandExplore,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
