package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocation(pageURL string) (LocationName, error) {
	url := baseURL + "/location-area"

	if pageURL != "" {
		url = pageURL
	}

	var locationListResp LocationName

	// Check if cache has stored the URL
	val, exist := c.Cache.Get(url)
	// If exist
	if exist {
		fmt.Println("Cache hit")
		err := json.Unmarshal(val, &locationListResp)
		if err != nil {

			return LocationName{}, err
		}
		return locationListResp, nil
	}

	//If not exist
	fmt.Println("Cache missed")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationName{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationName{}, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationName{}, err
	}

	// Unmarshal Json to struct
	err = json.Unmarshal([]byte(resBody), &locationListResp)
	if err != nil {
		return LocationName{}, err

	}

	c.Cache.Add(url, resBody)

	return locationListResp, nil
}
