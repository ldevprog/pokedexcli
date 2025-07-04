package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/levon-dalakyan/pokedexcli/internal/cli"
	"github.com/levon-dalakyan/pokedexcli/internal/pokecache"
	"github.com/levon-dalakyan/pokedexcli/internal/pokedata"
)

func CommandMap(
	cliCommands cli.CmdMap,
	config *pokedata.LocationAreasData,
	cache *pokecache.Cache,
) error {
	url := config.Next
	fetchAndUpdate(url, config, cache)
	return nil
}

func CommandMapB(
	cliCommands cli.CmdMap,
	config *pokedata.LocationAreasData,
	cache *pokecache.Cache,
) error {
	url := config.Previous
	if url == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	fetchAndUpdate(*url, config, cache)
	return nil
}

func fetchAndUpdate(
	url string,
	config *pokedata.LocationAreasData,
	cache *pokecache.Cache,
) error {
	var data []byte
	data, found := cache.Get(url)
	if found {
		fmt.Println("✅ data retrieved from cache!")
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		data, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
		}
		if err != nil {
			return err
		}
		cache.Add(url, data)
	}
	if err := json.Unmarshal(data, config); err != nil {
		return err
	}
	for _, la := range config.Result {
		fmt.Println(la.Name)
	}
	return nil
}
