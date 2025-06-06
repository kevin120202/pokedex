// Package main contains the command implementations for the Pokedex
package main

import (
	"fmt"
	"os"
)

// commandExit handles the "exit" command to terminate the Pokedex application
// It prints a goodbye message and exits the program with status code 0
func commandExit(cfg *config) error {
	// Print goodbye message to the user
	fmt.Print("Closing the Pokedex... Goodbye!")

	// Exit the program with status code 0 (success)
	// This immediately terminates the program
	os.Exit(0)

	// This line is never reached due to os.Exit(0)
	// It's included to satisfy the function signature
	return nil
}
