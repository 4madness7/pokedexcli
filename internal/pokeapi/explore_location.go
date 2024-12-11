package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/4madness7/pokedexcli/internal/pokecache"
)

func (c *Client) ExploreLocationArea(location string, cache *pokecache.Cache) (ExploredLocationArea, error) {
	url := baseURL + "/location-area/" + location

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

	cache.Add(location, data)

	return exploredLocationsResp, nil
}
