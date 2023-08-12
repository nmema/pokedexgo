package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (locationAreas, error) {
	endpont := "/location-area"
	fullURL := baseURL + endpont

	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		location := locationAreas{}
		err := json.Unmarshal(data, &location)
		if err != nil {
			return locationAreas{}, err
		}

		return location, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return locationAreas{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreas{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return locationAreas{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return locationAreas{}, err
	}

	location := locationAreas{}
	err = json.Unmarshal(data, &location)
	if err != nil {
		return locationAreas{}, err
	}

	c.cache.Add(fullURL, data)

	return location, nil
}

func (c *Client) PokemonByLocationArea(locationName string) ([]PokemonEncounters, error) {
	endpont := "/location-area"
	fullURL := baseURL + endpont + "/" + locationName

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		locationDetails := locationAreaDetails{}
		err := json.Unmarshal(data, &locationDetails)
		if err != nil {
			return []PokemonEncounters{}, err
		}

		return locationDetails.PokemonEncounters, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return []PokemonEncounters{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []PokemonEncounters{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return []PokemonEncounters{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return []PokemonEncounters{}, err
	}

	locationDetails := locationAreaDetails{}
	err = json.Unmarshal(data, &locationDetails)
	if err != nil {
		return []PokemonEncounters{}, err
	}

	c.cache.Add(fullURL, data)

	return locationDetails.PokemonEncounters, nil
}
