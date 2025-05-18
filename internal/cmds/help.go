package cmds

import (
	"fmt"
	"bootdev/pokedex/internal/pokeapi"
)

type HelpCommand struct {
	Cfg			*Config
	Clt			*pokeapi.Client
}

func NewHelpCommand(cfg *Config, clt *pokeapi.Client) CommandHandler {
	return HelpCommand{
		Cfg:			cfg,
		Clt:			clt,
	}
}

func (h HelpCommand) Name() string { return "help" }
func (h HelpCommand) Description() string { return "Displays a help message" }

func (h HelpCommand) Run(args []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, c := range InitializeCmdRegistry(h.Cfg, h.Clt) {
		fmt.Printf("%s: %s\n", c.Name(), c.Description())
	}
	return nil
}