package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Navigation struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

type Location struct {
	Name string `json:"name"`
}

type LocationArea struct {
	//Navigation Navigation
	Next      *string    `json:"next"`
	Previous  *string    `json:"previous"`
	Locations []Location `json:"results"`
}

func GetLocationAreas(url *string) (LocationArea, error) {
	var loc LocationArea

	res, err := http.Get(*url)
	if err != nil {
		return loc, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return loc, err
	}

	json.Unmarshal([]byte(data), &loc)

	return loc, nil
}
