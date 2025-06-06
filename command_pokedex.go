package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("your pokedex is empty")
	}
	fmt.Println("Your pokedex:")
	for key := range cfg.pokedex {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}
