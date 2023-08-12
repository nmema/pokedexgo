package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide one pokemon")
	}

	pokemonName := args[0]

	_, ok := conf.pokemonCaught[pokemonName]
	if ok {
		return errors.New("pokemon already caught")
	}

	pokemon, err := conf.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	fmt.Printf("%s was caught!\n", pokemonName)

	conf.pokemonCaught[pokemonName] = pokemon

	return nil
}
