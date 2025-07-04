package pokedata

type FetchConfig struct {
	Locations FetchLocationsData
	Location  FetchLocationData
}

type FetchLocationsData struct {
	Next     string   `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Entity `json:"results"`
}

type FetchLocationData struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Entity `json:"pokemon"`
}

type Entity struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
