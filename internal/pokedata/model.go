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
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Types          []PokemonType `json:"types"`
	Stats          []PokemonStat `json:"stats"`
}

type PokemonType struct {
	Slot int    `json:"slot"`
	Type Entity `json:"type"`
}

type PokemonStat struct {
	BaseStat int    `json:"base_stat"`
	Effort   int    `json:"effort"`
	Stat     Entity `json:"stat"`
}
