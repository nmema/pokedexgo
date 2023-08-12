package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (PokemonInformation, error) {
	endpont := "/pokemon"
	fullURL := baseURL + endpont + "/" + pokemonName

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		pokemonInfo := PokemonInformation{}
		err := json.Unmarshal(data, &pokemonInfo)
		if err != nil {
			return PokemonInformation{}, err
		}

		return pokemonInfo, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonInformation{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInformation{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return PokemonInformation{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInformation{}, err
	}

	pokemonInfo := PokemonInformation{}
	err = json.Unmarshal(data, &pokemonInfo)
	if err != nil {
		return PokemonInformation{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemonInfo, nil

}
