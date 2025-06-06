// Package main contains the core functionality of the Pokedex application
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kevin120202/pokedex/internal/pokeapi"
)

// config holds the application's state and dependencies
// It's passed to commands to maintain state between command executions
type config struct {
	// pokeapiClient is used to make HTTP requests to the PokeAPI
	pokeapiClient pokeapi.Client
	// nextLocationsURL stores the URL for the next page of locations
	nextLocationsURL *string
	// prevLocationsURL stores the URL for the previous page of locations
	prevLocationsURL *string
}

// startRepl initializes and runs the Read-Eval-Print Loop (REPL)
// It continuously reads user input, processes it, and executes commands
func startRepl(cfg *config) {
	// Create a new scanner to read from standard input (keyboard)
	// This allows us to read user input line by line
	reader := bufio.NewScanner(os.Stdin)

	// Main REPL loop - continues until the program is terminated
	for {
		// Display the prompt and wait for user input
		fmt.Print("Pokedex > ")

		// Read a line of input from the user
		reader.Scan()

		// Process the input: convert to lowercase and split into words
		words := cleanInput(reader.Text())
		// Skip to next iteration if no input was provided
		if len(words) == 0 {
			continue
		}

		// Get the first word as the command name
		commandName := words[0]
		// Look up the command in the registered commands map
		command, exists := getCommands()[commandName]

		if exists {
			// Execute the command with the config
			err := command.callback(cfg)
			if err != nil {
				// Print any errors that occur during command execution
				fmt.Println(err)
			}
			continue
		} else {
			// Handle unknown commands
			fmt.Println("Unknown command")
			continue
		}
	}
}

// cleanInput processes the input string by:
// 1. Converting it to lowercase
// 2. Splitting it into words (handling multiple spaces)
// Returns a slice of strings containing the processed words
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

// cliCommand represents a command that can be executed in the REPL
// Each command has a name, description, and callback function
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// getCommands returns a map of all available commands
// The map keys are command names, and the values are cliCommand structs
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
