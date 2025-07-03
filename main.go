package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(map[string]cliCommand) error
}

func main() {
	cliCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		inputCommand := cleanedInput[0]
		command, ok := cliCommands[inputCommand]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cliCommands)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func commandExit(cliCommands map[string]cliCommand) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cliCommands map[string]cliCommand) error {
	commandsHelp := ""
	for _, commandData := range cliCommands {
		commandsHelp += fmt.Sprintf("%s: %s\n", commandData.name, commandData.description)
	}
	fmt.Printf(`Welcome to the Pokedex!
Usage:

%s`, commandsHelp)
	return nil
}

func cleanInput(text string) []string {
	textTrimed := strings.TrimSpace(text)
	textLower := strings.ToLower(textTrimed)
	words := strings.Fields(textLower)
	return words
}
