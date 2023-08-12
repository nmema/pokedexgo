package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func commandHelp() {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	commands := getCommands()

	for _, command := range commands {
		fmt.Printf("%s:  %s", command.name, command.description)
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

var commands map[string]cliCommand = getCommands()

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	scanner.Scan()
	text := scanner.Text()

	fmt.Println(text)
	commands[text].callback()

	// commands[text].callback()
}
