package cli

import (
	"strings"

	"github.com/levon-dalakyan/pokedexcli/internal/pokecache"
	"github.com/levon-dalakyan/pokedexcli/internal/pokedata"
)

type CmdMap map[string]CliCommand

type CliCommand struct {
	Name        string
	Description string
	Callback    func(
		CmdMap,
		*pokedata.AppData,
		*pokecache.Cache,
		string,
	) error
}

func CleanInput(text string) []string {
	textTrimed := strings.TrimSpace(text)
	textLower := strings.ToLower(textTrimed)
	words := strings.Fields(textLower)
	return words
}
