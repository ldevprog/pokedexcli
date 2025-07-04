package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/levon-dalakyan/pokedexcli/internal/cli"
	"github.com/levon-dalakyan/pokedexcli/internal/pokecache"
	"github.com/levon-dalakyan/pokedexcli/internal/pokedata"
)

func main() {
	cliCommands := initCommands()
	config := initConfig()
	cache := pokecache.NewCache(10 * time.Second)

	runCli(cliCommands, &config, cache)
}

func runCli(
	cliCommands cli.CmdMap,
	config *pokedata.LocationAreasData,
	cache *pokecache.Cache,
) {
	prompt := "Pokedex >"
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cli.CleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}
		inputCommand := cleanedInput[0]
		command, ok := cliCommands[inputCommand]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.Callback(cliCommands, config, cache)
		if err != nil {
			fmt.Println(err)
		}
	}
}
