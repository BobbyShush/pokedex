package cmds

import (
	"fmt"
	"bootdev/pokedex/internal/pokeapi"
)

type ExploreCommand struct {
	Cfg			*Config
	Clt			*pokeapi.Client
}

func NewExploreCommand(cfg *Config, clt *pokeapi.Client) CommandHandler {
	return ExploreCommand{
		Cfg:			cfg,
		Clt:			clt,
	}
}

func (e ExploreCommand) Name() string { return "explore" }
func (e ExploreCommand) Description() string { return "Displays a list of all Pokemon within an area (1 argument: area name - use map cmd to list areas)"}

func (e ExploreCommand) Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Expected argument: area")
	}
	
	area, err := e.Clt.GetArea(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", args[0])
	fmt.Println("Found Pokemon")
	for _, pkmn := range area.PokemonEncounters {
		fmt.Printf(" - %s\n", pkmn.Pokemon.Name)
	}

	return nil
}