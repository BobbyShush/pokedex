package cmds

import (
	"bootdev/pokedex/internal/pokeapi"
)

type Config struct {
	Next		string
	Previous	string
}

type CliCommand struct {
	Name		string
	Description	string
	Callback	func(*Config, *pokeapi.Client) error
}

func InitializeCmdRegistry() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:			"exit",
			Description:	"Exit the Pokedex",
			Callback:		commandExit,
		},
		"help": {
			Name:			"help",
			Description:	"Displays a help message",
			Callback:		commandHelp,
		},
		"map": {
			Name:			"map",
			Description:	"Displays the names of 20 location areas in the Pokemon world",
			Callback:		commandMap,
		},
		"mapb": {
			Name:			"mapb",
			Description:	"Displays the previous page of location areas",
			Callback:		commandMapb,
		},
		"explore": {
			Name:			"explore",
			Description:	"Displays a list of all pokemons within a specified area",
			Callback:		commandExplore,
		},
	}
}