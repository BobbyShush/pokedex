package cmds

import (
	"fmt"
	"bootdev/pokedex/internal/pokeapi"
)

func commandHelp(cfg *Config, client *pokeapi.Client) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, c := range InitializeCmdRegistry() {
		fmt.Printf("%s: %s\n", c.Name, c.Description)
	}
	return nil
}