package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"bootdev/pokedex/internal/pokeapi"
	"bootdev/pokedex/internal/cmds"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commandRegistry := cmds.InitializeCmdRegistry()
	cfg := &cmds.Config{}
	clt := pokeapi.NewClient()
	
	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			fmt.Printf("scanner error: %v\n", scanner.Err())
    		break
		}

		text := cleanInput(scanner.Text())
		if len(text) < 1 {
			continue
		}

		command, exists := commandRegistry[text[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.Callback(cfg, clt); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
