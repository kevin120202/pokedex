package main

import (
	"time"

	"github.com/kevin120202/pokedex/internal/pokeapi"
	"github.com/kevin120202/pokedex/internal/pokecache"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	interval := 5 * time.Second
	cache := pokecache.NewCache(interval)

	startRepl(&cfg, cache)
}
