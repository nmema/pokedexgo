package main

import (
	"errors"
	"fmt"
)

func commandInspect(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide one pokemon")
	}

	pokemonName := args[0]

	information, ok := conf.pokemonCaught[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", information.Name)
	fmt.Printf("Height: %v\n", information.Height)
	fmt.Printf("Weight: %v\n", information.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range information.Stats {
		fmt.Printf("  - %s:  %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, typeInfo := range information.Types {
		fmt.Printf("  - %s\n", typeInfo.Type.Name)
	}

	return nil
}
