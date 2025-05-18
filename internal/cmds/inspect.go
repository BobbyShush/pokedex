package cmds

import (
	"fmt"
	"bootdev/pokedex/internal/pokeapi"
)

type InspectCommand struct {
	Clt			*pokeapi.Client
}

func NewInspectCommand(clt *pokeapi.Client) CommandHandler {
	return InspectCommand{
		Clt:			clt,
	}
}

func (i InspectCommand) Name() string { return "inspect" }
func (i InspectCommand) Description() string { return "Display the information of a Pokemon registered in the Pokedex (1 arg: Pokemon name)"}

func (i InspectCommand) Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Expected argument: Pokemon name")
	}
	pkmnName := args[0]

	pkmn, exists := i.Clt.Pokedex[pkmnName]
	if !exists {
		return fmt.Errorf("%s isn't registered yet. Catch it to register it. Is its name spelled right?", pkmnName)
	}

	fmt.Printf("Name: %s\n", pkmnName)
	fmt.Printf("Height: %d\n", pkmn.Height)
	fmt.Printf("Weight: %d\n", pkmn.Weight)

	fmt.Println("Stats:")
	for _, stat := range pkmn.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pkmn.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}