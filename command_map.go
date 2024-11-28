package main

import (
	"errors"
	"fmt"

	"github.com/kevin120202/pokedex/internal/pokeapi"
)

func mapCallback(config *config) error {
	err := mapApiCall(config, config.nextLocationAreaUrl)
	if err != nil {
		return err
	}

	return nil
}

func mapbCallback(config *config) error {
	if config.prevLocationAreaUrl == nil {
		return errors.New("page unavailable")
	}

	err := mapApiCall(config, config.prevLocationAreaUrl)
	if err != nil {
		return err
	}

	return nil
}

func mapApiCall(config *config, locationUrl *string) error {
	pokeapiClient := pokeapi.NewClient()
	res, err := pokeapiClient.ListLocationAreas(locationUrl)
	if err != nil {
		return err
	}

	config.nextLocationAreaUrl = res.Next
	config.prevLocationAreaUrl = res.Previous

	for i := 0; i < len(res.Results); i++ {
		fmt.Printf(" - %s\n", res.Results[i].Name)
	}

	return nil
}
