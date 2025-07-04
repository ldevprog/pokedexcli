package commands

import (
	"fmt"
	"os"
	"sort"

	"github.com/levon-dalakyan/pokedexcli/internal/cli"
	"github.com/levon-dalakyan/pokedexcli/internal/pokecache"
	"github.com/levon-dalakyan/pokedexcli/internal/pokedata"
)

func CommandExit(
	cliCommands cli.CmdMap,
	config *pokedata.FetchConfig,
	cache *pokecache.Cache,
	argument string,
) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(
	cliCommands cli.CmdMap,
	config *pokedata.FetchConfig,
	cache *pokecache.Cache,
	argument string,
) error {
	var commands []cli.CliCommand
	for _, cmd := range cliCommands {
		commands = append(commands, cmd)
	}
	sort.Slice(commands, func(i, j int) bool {
		return commands[i].Name < commands[j].Name
	})
	commandsOut := ""
	for _, cmd := range commands {
		commandsOut += fmt.Sprintf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Printf(`Welcome to the Pokedex!
Usage:

%s`, commandsOut)
	return nil
}
