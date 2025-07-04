package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/levon-dalakyan/pokedexcli/internal/cli"
	"github.com/levon-dalakyan/pokedexcli/internal/pokecache"
	"github.com/levon-dalakyan/pokedexcli/internal/pokedata"
)

func CommandMap(
	cliCommands cli.CmdMap,
	config *pokedata.AppData,
	cache *pokecache.Cache,
	argument string,
) error {
	url := config.Locations.Next

	err := mapAndPrint(url, config, cache)
	return err
}

func CommandMapB(
	cliCommands cli.CmdMap,
	config *pokedata.AppData,
	cache *pokecache.Cache,
	argument string,
) error {
	url := config.Locations.Previous
	if url == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	err := mapAndPrint(*url, config, cache)
	return err
}

func CommandExplore(
	cliCommands cli.CmdMap,
	config *pokedata.AppData,
	cache *pokecache.Cache,
	argument string,
) error {
	if argument == "" {
		fmt.Println("Command \"explore\" needs an argument of area name")
		return nil
	}
	fmt.Printf("Exploring %s...\n", argument)

	url := "https://pokeapi.co/api/v2/location-area/" + argument
	data, err := fetchAndMemoize(url, cache)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &config.Location); err != nil {
		return err
	}

	for _, encounter := range config.Location.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func CommandCatch(
	cliCommands cli.CmdMap,
	config *pokedata.AppData,
	cache *pokecache.Cache,
	argument string,
) error {
	if argument == "" {
		fmt.Println("Command \"catch\" needs an argument of pokemon name")
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", argument)

	url := "https://pokeapi.co/api/v2/pokemon/" + argument
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &config.Pokemon); err != nil {
		return err
	}

	pokemon := config.Pokemon
	catchChance := 1.0 - (float64(pokemon.BaseExperience-50) / 100)
	roll := rand.Float64()
	if roll < catchChance {
		config.Pokedex[pokemon.Name] = pokemon
		fmt.Printf("🎉 you caught %s!\n", pokemon.Name)
	} else {
		fmt.Printf("💨 %s escaped!\n", pokemon.Name)
	}

	return nil
}

func mapAndPrint(url string, config *pokedata.AppData, cache *pokecache.Cache) error {
	data, err := fetchAndMemoize(url, cache)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &config.Locations); err != nil {
		return err
	}

	for _, la := range config.Locations.Results {
		fmt.Println(la.Name)
	}

	return nil
}

func fetchAndMemoize(url string, cache *pokecache.Cache) ([]byte, error) {
	var data []byte
	data, found := cache.Get(url)

	if found {
		fmt.Println("✅ data retrieved from cache!")
	} else {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return nil, fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		cache.Add(url, data)
	}

	return data, nil
}
