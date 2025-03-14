package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationName struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocation(pageURL string) (LocationName, error) {
	url := baseURL + "/location-area"

	if pageURL != "" {
		url = pageURL
	}
	var resource LocationName
	resp, err := http.Get(url)
	if err != nil {
		return resource, fmt.Errorf("error in fetching pokemon location: %w", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resource, fmt.Errorf("error in reading body:%w", err)
	}

	// Unmarshal Json to struct
	err = json.Unmarshal([]byte(body), &resource)
	if err != nil {
		return resource, fmt.Errorf("error in unmarshall locations from json format")

	}

	return resource, nil
}
