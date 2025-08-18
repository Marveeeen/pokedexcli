package pokeapi

import (
	"net/http"
	"fmt"
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
			return RespShallowLocations{}, fmt.Errorf("failed to unmarshal cached data: %v", err)
		}

		return cachedResp, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return RespShallowLocations{}, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return RespShallowLocations{}, fmt.Errorf("failed to fetch locations, status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("failed to read response body: %v", err)
	}

	c.cache.Add(url, data)

	locationsResp := RespShallowLocations{}
	if err := json.Unmarshal(data, &locationsResp); err != nil {
		return RespShallowLocations{}, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return locationsResp, nil
}