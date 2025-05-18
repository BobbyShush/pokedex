package cmds

import (
	"fmt"
	"bootdev/pokedex/internal/pokeapi"
)

type MapCommand struct {
	Cfg			*Config
	Clt			*pokeapi.Client
}

func NewMapCommand(cfg *Config, clt *pokeapi.Client) CommandHandler {
	return MapCommand{
		Cfg:			cfg,
		Clt:			clt,
	}
}

func (m MapCommand) Name() string { return "map" }
func (m MapCommand) Description() string { return "Displays the names of 20 location areas in the Pokemon world" }

func (m MapCommand) Run(args []string) error {
	url := ""
	if m.Cfg.Next != "" {
		url = m.Cfg.Next
	}
	return displayLocationAreas(url, m.Cfg, m.Clt)	
}

func displayLocationAreas(url string, cfg *Config, client *pokeapi.Client) error {
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