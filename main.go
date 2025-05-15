package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

type cliCommand struct {
	name		string
	description	string
	callback	func() error
	config		struct {
		Next		string
		Previous	string
	}
}

var commandRegistry map[string]cliCommand

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commandRegistry = initializeCmdRegistry()
	
	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			fmt.Printf("scanner error: %v\n", scanner.Err())
    		break
		}

		text := cleanInput(scanner.Text())
		if len(text) < 1 {
			continue
		}

		command, exists := commandRegistry[text[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func initializeCmdRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback:		commandExit,
		},
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		commandHelp,
		},
		"map": {
			name:			"map",
			description:	"Displays the names of 20 location areas in the Pokemon world",
			callback:		commandMap,
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, c := range commandRegistry {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap() error {
	return nil
}