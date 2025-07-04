package pokedata

type AppData struct {
	Locations FetchLocationsData
	Location  FetchLocationData
	Pokemon   Pokemon
	Pokedex   map[string]Pokemon
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

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}
