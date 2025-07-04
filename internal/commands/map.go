package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/levon-dalakyan/pokedexcli/internal/cli"
	"github.com/levon-dalakyan/pokedexcli/internal/pokedata"
)

func CommandMap(cliCommands cli.CmdMap, config *pokedata.LocationAreasData) error {
	url := config.Next
	fetchAndUpdate(url, config)
	return nil
}

func CommandMapB(cliCommands cli.CmdMap, config *pokedata.LocationAreasData) error {
	url := config.Previous
	if url == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	fetchAndUpdate(*url, config)
	return nil
}

func fetchAndUpdate(url string, config *pokedata.LocationAreasData) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
	}
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, config); err != nil {
		return err
	}
	for _, la := range config.Result {
		fmt.Println(la.Name)
	}
	return nil
}
