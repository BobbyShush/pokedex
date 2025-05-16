package cmds

import (
	"fmt"
	"os"
	"bootdev/pokedex/internal/pokeapi"
)

func commandExit(cfg *Config, client *pokeapi.Client) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}