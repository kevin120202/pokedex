package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// type config struct {
// 	Next string
// 	Prev string
// }

// startRepl initializes and runs the Read-Eval-Print Loop (REPL)
// It continuously reads user input, processes it, and executes commands
func startRepl() {
	// Create a new scanner to read from standard input (keyboard)
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		// Read a line of input from the user
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 locations areas",
			callback:    commandMap,
		},
	}
}
