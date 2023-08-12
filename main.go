package main

import (
	"time"

	"github.com/nmema/pokedexgo/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	pokemonCaught map[string]pokeapi.PokemonInformation
}

func main() {
	config := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		pokemonCaught: make(map[string]pokeapi.PokemonInformation),
	}

	startRelp(&config)
}
