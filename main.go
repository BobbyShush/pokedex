package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"bootdev/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name		string
	description	string
	callback	func(*config) error
}

type config struct {
	Next		string
	Previous	string
}

var commandRegistry map[string]cliCommand

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commandRegistry = initializeCmdRegistry()
	cfg := &config{}
	
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
		if err := command.callback(cfg); err != nil {
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
		"mapb": {
			name:			"mapb",
			description:	"Displays the previous page of location areas",
			callback:		commandMapb,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, c := range commandRegistry {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	url := ""
	if cfg.Next != "" {
		url = cfg.Next
	}
	return displayLocationAreas(url, cfg)	
}

func commandMapb(cfg *config) error {
	if cfg.Previous == "" {
		return fmt.Errorf("you're on the first page")
	}
	return displayLocationAreas(cfg.Previous, cfg)
}

func displayLocationAreas(url string, cfg *config) error {
	client := pokeapi.NewClient()
	locationAreas, err := client.GetLocationAreas(url)
	if err != nil {
		return err
	}

	cfg.Previous = locationAreas.Previous
	cfg.Next = locationAreas.Next

	for _, la := range locationAreas.Results {
		fmt.Println(la.Name)
	}

	return nil
}