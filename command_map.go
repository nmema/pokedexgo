package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type locationArea struct {
	Count    int `json:"count"`
	Next     any `json:"next"`
	Previous any `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
	Next     any
	Previous any
}

var initialConfig = config{
	Next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	Previous: nil,
}

func commandMap(conf *config) error {
	str, ok := conf.Next.(string)

	if !ok {
		return fmt.Errorf("you are on the last page")
	}

	res, err := http.Get(str)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf(
			"Response failed with:\n - status code: %d\n - body: %s\n",
			res.StatusCode,
			body,
		)
	}

	if err != nil {
		log.Fatal(err)
	}

	location := locationArea{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		fmt.Println(err)
	}

	for _, result := range location.Results {
		fmt.Println(result.Name)
	}

	conf.Next = location.Next
	conf.Previous = location.Previous

	return nil
}

func commandMapb(conf *config) error {

	str, ok := conf.Previous.(string)

	if !ok {
		return fmt.Errorf("you are on the first page")
	}

	res, err := http.Get(str)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf(
			"Response failed with:\n - status code: %d\n - body: %s\n",
			res.StatusCode,
			body,
		)
	}

	if err != nil {
		log.Fatal(err)
	}

	location := locationArea{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		fmt.Println(err)
	}

	for _, result := range location.Results {
		fmt.Println(result.Name)
	}

	conf.Next = location.Next
	conf.Previous = location.Previous

	return nil
}
