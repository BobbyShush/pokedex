package cmds

import (
	"fmt"
	"math"
	"math/rand"
	"bootdev/pokedex/internal/pokeapi"
)

type CatchCommand struct{
	Clt			*pokeapi.Client
}

func NewCatchCommand(clt *pokeapi.Client) CommandHandler {
	return CatchCommand{
		Clt:			clt,
	}
}

func (c CatchCommand) Name() string { return "catch" }
func (c CatchCommand) Description() string { return "Attempts to catch a Pokemon (1 arg: Pokemon name)"}

func (c CatchCommand) Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Expected argument: Pokemon name")
	}

	pkmnName := args[0]
	pkmn, err := c.Clt.GetPokemon(pkmnName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pkmnName)
	if isCaught(pkmn) {
		fmt.Printf("%s was caught!\n", pkmnName)
		c.Clt.Pokedex[pkmnName] = pkmn
		return nil
	}
	fmt.Printf("%s escaped!\n", pkmnName)
	return nil
}

func isCaught(pkmn pokeapi.PokemonResp) bool {
	const MEWTWOBASEXP = 340.0
	scale := 100.0 / MEWTWOBASEXP
	randNum := rand.Intn(100) + 1
	catchRate := math.Max(5.0, 100.0 - (float64(pkmn.BaseExperience) * scale))
	if float64(randNum) <= catchRate {
		return true
	}
	return false
}