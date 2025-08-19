package pokeapi

import (
	"net/http"
	"io"
	"encoding/json"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if cachedData, ok := c.cache.Get(url); ok {
		var cachedResp Location
		if err := json.Unmarshal(cachedData, &cachedResp); err != nil {
			return Location{}, err
		}

		return cachedResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	if err := json.Unmarshal(data, &locationResp); err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)
	return locationResp, nil
}