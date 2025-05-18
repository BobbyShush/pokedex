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
	cfg := &cmds.Config{}
	clt := pokeapi.NewClient()
	commandRegistry := cmds.InitializeCmdRegistry(cfg, clt)

	
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

		cmdName := text[0]
		args := []string{}
		if len(text) > 1 {
			args = text[1:]
		}
		command, exists := commandRegistry[cmdName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.Run(args); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
