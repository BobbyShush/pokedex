package cmds

import (
	"fmt"
	"bootdev/pokedex/internal/pokeapi"
)

type PokedexCommand struct {
	Clt			*pokeapi.Client
}

func NewPokedexCommand(clt *pokeapi.Client) CommandHandler {
	return PokedexCommand{
		Clt:			clt,
	}
}

func (p PokedexCommand) Name() string { return "pokedex" }
func (p PokedexCommand) Description() string { return "Lists all registered Pokemon" }

func (p PokedexCommand) Run(args []string) error {
	if len(p.Clt.Pokedex) < 1 {
		return fmt.Errorf("No Pokemon is currently registered in the Pokedex")
	}

	fmt.Println("Your Pokedex:")
	for pkmn := range p.Clt.Pokedex {
		fmt.Printf(" - %s\n", pkmn)
	}
	return nil
}