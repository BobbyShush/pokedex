package cmds

import (
	"fmt"
	"bootdev/pokedex/internal/pokeapi"
)

type MapbCommand struct {
	Cfg			*Config
	Clt			*pokeapi.Client
}

func NewMapbCommand(cfg *Config, clt *pokeapi.Client) CommandHandler {
	return MapbCommand{
		Cfg:			cfg,
		Clt:			clt,
	}
}

func (m MapbCommand) Name() string { return "mapb" }
func (m MapbCommand) Description() string { return "Displays the previous page of location areas"}

func (m MapbCommand) Run(args []string) error {
	if m.Cfg.Previous == "" {
		return fmt.Errorf("you're on the first page")
	}
	return displayLocationAreas(m.Cfg.Previous, m.Cfg, m.Clt)
}