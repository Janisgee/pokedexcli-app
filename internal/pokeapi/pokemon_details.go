package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonDetails(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon" + "/" + pokemonName

	var pokemonResp Pokemon

	// Check if cache has stored the URL
	val, exist := c.Cache.Get(url)
	// If exist
	if exist {
		fmt.Println("Cache hit")
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {

			return Pokemon{}, err
		}
		return pokemonResp, nil
	}
	//If not exist
	fmt.Println("Cache missed")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// Unmarshal Json to struct
	err = json.Unmarshal([]byte(resBody), &pokemonResp)
	if err != nil {
		return Pokemon{}, err

	}

	c.Cache.Add(url, resBody)

	return pokemonResp, nil
}
