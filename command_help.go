package main

import "fmt"

func commandHelp() error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	commands := getCommands()

	for _, command := range commands {
		fmt.Printf(" - %s:  %s\n", command.name, command.description)
	}

	return nil
}
