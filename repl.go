package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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

		val.callback()
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func()
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
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
