package main

import (
	"github.com/levon-dalakyan/pokedexcli/internal/cli"
	"github.com/levon-dalakyan/pokedexcli/internal/commands"
	"github.com/levon-dalakyan/pokedexcli/internal/pokedata"
)

func initCommands() cli.CmdMap {
	return cli.CmdMap{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commands.CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commands.CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays 20 NEXT Poke Location Areas",
			Callback:    commands.CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays 20 PREV Poke Location Areas",
			Callback:    commands.CommandMapB,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays list of Pokémon located on area provided as argument",
			Callback:    commands.CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Throws a Pokeball at pokemon to catch. Provide pokemon name as an argument",
			Callback:    commands.CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect already caught pokemon. Use name as an argument",
			Callback:    commands.CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Display all caught pokemons",
			Callback:    commands.CommandPokedex,
		},
	}
}

func initConfig() pokedata.AppData {
	return pokedata.AppData{
		Locations: pokedata.FetchLocationsData{
			Next:     "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
			Previous: nil,
		},
		Pokedex: make(map[string]pokedata.Pokemon),
	}
}
