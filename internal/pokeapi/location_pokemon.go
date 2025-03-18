package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationPokemon(locationName string) (LocationDetails, error) {

	url := baseURL + "/location-area" + "/" + locationName

	var locationDetailResp LocationDetails

	// Check if cache has stored the URL
	val, exist := c.Cache.Get(url)
	// If exist
	if exist {
		fmt.Println("Cache hit")
		err := json.Unmarshal(val, &locationDetailResp)
		if err != nil {

			return LocationDetails{}, err
		}
		return locationDetailResp, nil
	}
	//If not exist
	fmt.Println("Cache missed")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetails{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationDetails{}, err
	}

	// Unmarshal Json to struct
	err = json.Unmarshal([]byte(resBody), &locationDetailResp)
	if err != nil {
		return LocationDetails{}, err

	}

	c.Cache.Add(url, resBody)

	return locationDetailResp, nil
}
