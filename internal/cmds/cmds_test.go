package cmds

import (
	"testing"
	"bytes"
	"os"
	"io"
	"fmt"
	"strings"
	"bootdev/pokedex/internal/pokeapi"
)

func TestInspect(t *testing.T) {
	cases := []struct {
		name string
		errStr string
		outStr string
	}{
		{
			name: "invalid",
			errStr: "invalid isn't registered yet. Catch it to register it. Is its name spelled right?",
		},
		{
			name: "lugia",
			errStr: "lugia isn't registered yet. Catch it to register it. Is its name spelled right?",
		},
		{
			name: "pikachu",
			outStr: `Name: pikachu
Height: 4
Weight: 2
Stats:
  -hp: 1
  -attack: 2
  -defense: 3
  -special-attack: 4
  -special-defense: 5
  -speed: 6
Types:
  - electric
`,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			client := pokeapi.NewClient()
			inspectCommand := InspectCommand{ Clt: client }
			inspectCommand.Clt.Pokedex["pikachu"] = pokeapi.PokemonResp{
				Height: 4,
				Weight: 2,
				Stats: []struct{
					BaseStat int `json:"base_stat"`
					Effort   int `json:"effort"`
					Stat     struct {
						Name string `json:"name"`
						URL  string `json:"url"`
					} `json:"stat"`
				}{
					{
						BaseStat: 1,
						Effort: 0,
						Stat: struct {
							Name string `json:"name"`
							URL string	`json:"url"`
						}{
							Name: "hp",
							URL: "",
						},
					},
					{
						BaseStat: 2,
						Effort: 0,
						Stat: struct {
							Name string `json:"name"`
							URL string	`json:"url"`
						}{
							Name: "attack",
							URL: "",
						},
					},
					{
						BaseStat: 3,
						Effort: 0,
						Stat: struct {
							Name string `json:"name"`
							URL string	`json:"url"`
						}{
							Name: "defense",
							URL: "",
						},
					},
					{
						BaseStat: 4,
						Effort: 0,
						Stat: struct {
							Name string `json:"name"`
							URL string	`json:"url"`
						}{
							Name: "special-attack",
							URL: "",
						},
					},
					{
						BaseStat: 5,
						Effort: 0,
						Stat: struct {
							Name string `json:"name"`
							URL string	`json:"url"`
						}{
							Name: "special-defense",
							URL: "",
						},
					},
					{
						BaseStat: 6,
						Effort: 0,
						Stat: struct {
							Name string `json:"name"`
							URL string	`json:"url"`
						}{
							Name: "speed",
							URL: "",
						},
					},
				},
				Types: []struct{
					Slot int `json:"slot"`
					Type struct {
						Name string `json:"name"`
						URL  string `json:"url"`
					} `json:"type"`
				}{
					{
						Slot: 0,
						Type: struct{ 
							Name string `json:"name"`
							URL string `json:"url"`
						}{
							Name: "electric",
							URL: "",
						},
					},
				},
			}

			reader, writer, err := os.Pipe()
			if err != nil {
				t.Errorf("error reading stdout: %v", err)
			}

			origStdout := os.Stdout
			defer func(){ os.Stdout = origStdout }()
			os.Stdout = writer

			err = inspectCommand.Run([]string{c.name})
			writer.Close()

			var buf bytes.Buffer
			io.Copy(&buf, reader)
			output := buf.String()

			if err != nil {
				if err.Error() != c.errStr {
					t.Errorf("expected output: %s\nactual output: %s", c.errStr, err.Error())
					return
				}
				return
			}

			if strings.TrimSpace(output) != strings.TrimSpace(c.outStr) {
				t.Errorf("expected output:\n%q\n\nactual output:\n%q", c.outStr, output)
				return
			}
		})
	}
}