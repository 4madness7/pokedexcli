package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocationArea(location string) (ExploredLocationArea, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.Cache.Get(url); ok {
		location := ExploredLocationArea{}
		err := json.Unmarshal(val, &location)
		if err != nil {
			return ExploredLocationArea{}, err
		}
		return location, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ExploredLocationArea{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ExploredLocationArea{}, nil
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ExploredLocationArea{}, nil
	}

	exploredLocationsResp := ExploredLocationArea{}
	err = json.Unmarshal(data, &exploredLocationsResp)
	if err != nil {
		return ExploredLocationArea{}, nil
	}

	c.Cache.Add(url, data)

	return exploredLocationsResp, nil
}
