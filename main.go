package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		inputCleaned := cleanInput(input)
		firstWord := inputCleaned[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}

func cleanInput(text string) []string {
	textTrimed := strings.TrimSpace(text)
	textLower := strings.ToLower(textTrimed)
	words := strings.Fields(textLower)
	return words
}
