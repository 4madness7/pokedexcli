package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (DataLocationArea, error) {
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

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return DataLocationArea{}, err
	}

	locationsResp := DataLocationArea{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return DataLocationArea{}, err
	}

	return locationsResp, nil
}
