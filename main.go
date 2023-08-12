package main

import "github.com/nmema/pokedexgo/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

func main() {
	config := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRelp(&config)
}
