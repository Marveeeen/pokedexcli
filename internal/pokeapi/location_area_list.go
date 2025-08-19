package pokeapi

import (
	"net/http"
	"io"
	"encoding/json"
)

func (c *Client) ListLocations(pageUrl *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area/"
	if pageUrl != nil {
		url = *pageUrl
	}

	if cachedData, ok := c.cache.Get(url); ok {
		var cachedResp RespShallowLocations
		if err := json.Unmarshal(cachedData, &cachedResp); err != nil {
			return RespShallowLocations{}, err
		}

		return cachedResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	if err := json.Unmarshal(data, &locationsResp); err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, data)
	return locationsResp, nil
}