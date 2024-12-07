package main

import (
	"os"

	"github.com/kevin120202/pokedex/internal/pokecache"
)

func exitCallback(config *config, cache *pokecache.Cache) error {
	os.Exit(0)
	return nil
}
