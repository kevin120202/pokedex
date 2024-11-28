package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		commands := getCommands()
		command := cleaned[0]

		val, ok := commands[command]
		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := val.callback(config)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    helpCallback,
		},
		"exit": {
			name:        "exit",
			description: "Kills the terminal",
			callback:    exitCallback,
		},
		"map": {
			name:        "map",
			description: "Prints 20 locations",
			callback:    mapCallback,
		},
		"mapb": {
			name:        "mapb",
			description: "Prints previus 20 locations",
			callback:    mapbCallback,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
