package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	textTrimed := strings.TrimSpace(text)
	textLower := strings.ToLower(textTrimed)
	words := strings.Fields(textLower)
	return words
}
