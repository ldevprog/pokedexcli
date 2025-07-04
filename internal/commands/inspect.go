package commands

import (
	"fmt"

	"github.com/levon-dalakyan/pokedexcli/internal/cli"
	"github.com/levon-dalakyan/pokedexcli/internal/pokecache"
	"github.com/levon-dalakyan/pokedexcli/internal/pokedata"
)

func CommandInspect(
	cliCommands cli.CmdMap,
	config *pokedata.AppData,
	cache *pokecache.Cache,
	argument string,
) error {
	if argument == "" {
		fmt.Println("Command \"inspect\" needs an argument of pokemon name")
		return nil
	}

	pokemon, found := config.Pokedex[argument]
	if !found {
		fmt.Printf("you have not caught %s\n", argument)
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)

	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("\t- %s\n", t.Type.Name)
	}

	return nil
}
