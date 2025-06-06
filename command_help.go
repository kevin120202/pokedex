// Package main contains the command implementations for the Pokedex
package main

import "fmt"

// commandHelp handles the "help" command to display available commands and their descriptions
// It prints a welcome message and lists all registered commands with their descriptions
func commandHelp(cfg *config) error {
	// Print a blank line for better formatting
	fmt.Println()

	// Display welcome message
	fmt.Println("Welcome to the Pokedex!")

	// Print usage header
	fmt.Println("Usage:")
	fmt.Println()

	// Iterate through all registered commands and print their names and descriptions
	// getCommands() returns a map of all available commands
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	// Print a final blank line for better formatting
	fmt.Println()
	return nil
}
