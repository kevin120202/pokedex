// Package main is the entry point for the Pokedex application
package main

import (
	"time"

	"github.com/kevin120202/pokedex/internal/pokeapi"
)

// main is the entry point function that initializes and starts the Pokedex
func main() {
	// Create a new PokeAPI client with a 5-second timeout
	// This client will be used to make HTTP requests to the PokeAPI
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)

	// Initialize the config struct with the PokeAPI client
	// This config will be passed around to maintain state between commands
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	// Start the REPL (Read-Eval-Print Loop) with the config
	// This begins the interactive command-line interface
	startRepl(cfg)
}
