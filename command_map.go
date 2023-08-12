package main

import (
	"errors"
	"fmt"
)

func commandMap(conf *config) error {
	resp, err := conf.pokeapiClient.ListLocationAreas(conf.Next)
	if err != nil {
		return err
	}

	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}

	conf.Next = resp.Next
	conf.Previous = resp.Previous

	return nil
}

func commandMapb(conf *config) error {
	if conf.Previous == nil {
		return errors.New("you are on the first page")
	}

	resp, err := conf.pokeapiClient.ListLocationAreas(conf.Previous)
	if err != nil {
		return err
	}

	for _, result := range resp.Results {
		fmt.Println(result.Name)
	}

	conf.Next = resp.Next
	conf.Previous = resp.Previous

	return nil
}
