package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (DataLocationArea, error) {
	url := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.Cache.Get(url); ok {
		locationsResp := DataLocationArea{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return DataLocationArea{}, err
		}
		return locationsResp, nil
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

	c.Cache.Add(url, data)

	return locationsResp, nil
}
