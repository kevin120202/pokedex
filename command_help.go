package main

import (
	"fmt"

	"github.com/kevin120202/pokedex/internal/pokecache"
)

func helpCallback(config *config, cache *pokecache.Cache) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Here are your available commands")
	fmt.Println()

	commands := getCommands()

	for _, cm := range commands {
		fmt.Printf("%s: %s\n", cm.name, cm.description)
	}

	fmt.Println("")

	return nil
}
