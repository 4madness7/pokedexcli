package pokeapi

import (
	"encoding/json"
	"github.com/4madness7/pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string, cache *pokecache.Cache) (DataLocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return DataLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return DataLocationArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return DataLocationArea{}, err
	}

	locationsResp := DataLocationArea{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return DataLocationArea{}, err
	}

	cache.Add(url, data)

	return locationsResp, nil
}
