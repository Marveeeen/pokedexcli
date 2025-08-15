package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type paginate struct {
	Next string
	Previous any
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	config := &paginate{
		Next: "",
		Previous: nil,
	}
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command:", commandName)
			continue
		}

		err := command.callback(config)
		if err != nil {
			fmt.Println("Error executing command '%s': %v\n", commandName, err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *paginate) error
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
			description: "Display help information",
			callback:    commandHelp,
		},
	}
}
