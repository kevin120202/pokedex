package main

import "fmt"

func helpCallback(config *config) error {
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
