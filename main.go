package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/levon-dalakyan/pokedexcli/internal/cli"
	"github.com/levon-dalakyan/pokedexcli/internal/pokedata"
)

func main() {
	cliCommands := initCommands()
	config := initConfig()

	runCli(cliCommands, &config)
}

func runCli(cliCommands cli.CmdMap, config *pokedata.LocationAreasData) {
	scanner := bufio.NewScanner(os.Stdin)
	prompt := "Pokedex >"

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
		err := command.Callback(cliCommands, config)
		if err != nil {
			fmt.Println(err)
		}
	}
}
