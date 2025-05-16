package cmds

import (
	"fmt"
	"bootdev/pokedex/internal/pokeapi"
)

func commandMap(cfg *Config, client *pokeapi.Client) error {
	url := ""
	if cfg.Next != "" {
		url = cfg.Next
	}
	return displayLocationAreas(url, cfg, client)	
}

func commandMapb(cfg *Config, client *pokeapi.Client) error {
	if cfg.Previous == "" {
		return fmt.Errorf("you're on the first page")
	}
	return displayLocationAreas(cfg.Previous, cfg, client)
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