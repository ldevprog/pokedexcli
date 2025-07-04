package pokedata

type LocationAreasData struct {
	Next     string         `json:"next"`
	Previous *string        `json:"previous"`
	Result   []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
