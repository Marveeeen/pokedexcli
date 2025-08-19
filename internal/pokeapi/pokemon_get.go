package pokeapi

import (
	"net/http"
	"io"
	"encoding/json"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if cachedData, ok := c.cache.Get(url); ok {
		var cachedResp Pokemon
		if err := json.Unmarshal(cachedData, &cachedResp); err != nil {
			return Pokemon{}, err
		}

		return cachedResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()
	
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	if err := json.Unmarshal(data, &pokemonResp); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return pokemonResp, nil
}