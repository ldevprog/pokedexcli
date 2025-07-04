package cli

import (
	"strings"

	"github.com/levon-dalakyan/pokedexcli/internal/pokedata"
)

type CmdMap map[string]CliCommand

type CliCommand struct {
	Name        string
	Description string
	Callback    func(CmdMap, *pokedata.LocationAreasData) error
}

func CleanInput(text string) []string {
	textTrimed := strings.TrimSpace(text)
	textLower := strings.ToLower(textTrimed)
	words := strings.Fields(textLower)
	return words
}
