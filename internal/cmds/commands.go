package cmds

import (
	"bootdev/pokedex/internal/pokeapi"
)

type Config struct {
	Next		string
	Previous	string
}

type CommandHandler interface {
	Run(args []string) error
	Name() string
	Description() string
}

func InitializeCmdRegistry(cfg *Config, clt *pokeapi.Client) map[string]CommandHandler {
	return map[string]CommandHandler{
		"exit": NewExitCommand(),
		"help": NewHelpCommand(cfg, clt),
		"map":	NewMapCommand(cfg, clt),
		"mapb": NewMapbCommand(cfg, clt),
		"explore": NewExploreCommand(cfg, clt),
		"catch": NewCatchCommand(clt),
		"inspect": NewInspectCommand(clt),
		"pokedex": NewPokedexCommand(clt),
	}
}
