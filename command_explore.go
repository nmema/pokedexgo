package main

import (
	"errors"
	"fmt"
)

func commandExplore(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide one location area")
	}

	pkmEncounters, err := conf.pokeapiClient.PokemonByLocationArea(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Pokemons in that area:")
	for _, pokemon := range pkmEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
