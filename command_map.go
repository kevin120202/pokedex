package main

import (
	"errors"
	"fmt"

	"github.com/kevin120202/pokedex/internal/pokeapi"
	"github.com/kevin120202/pokedex/internal/pokecache"
)

func mapCallback(config *config, cache *pokecache.Cache) error {
	err := mapApiCall(config, config.nextLocationAreaUrl, cache)
	if err != nil {
		return err
	}

	return nil
}

func mapbCallback(config *config, cache *pokecache.Cache) error {
	if config.prevLocationAreaUrl == nil {
		return errors.New("page unavailable")
	}

	err := mapApiCall(config, config.prevLocationAreaUrl, cache)
	if err != nil {
		return err
	}

	return nil
}

func mapApiCall(config *config, locationUrl *string, cache *pokecache.Cache) error {
	pokeapiClient := pokeapi.NewClient()
	res, err := pokeapiClient.ListLocationAreas(locationUrl, cache)
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
