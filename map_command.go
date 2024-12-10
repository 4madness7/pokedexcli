package main

import (
    "io"
    "encoding/json"
    "fmt"
)

const (
	baseURL = "https://pokeapi.co/api/v2/"
)

type DataLocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}


func getLocationArea(path *string, cfg *config) (DataLocationArea, error) {
	res, err := cfg.Client.Get(*path)
	if err != nil {
		return DataLocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return DataLocationArea{}, err
	}

	var values DataLocationArea
	err = json.Unmarshal(data, &values)
	if err != nil {
		return DataLocationArea{}, err
	}

    cfg.Next = values.Next
    cfg.Previous = values.Previous

	return values, nil
}

func mapfCallback(cfg *config) error {
	fullPath := baseURL + "location-area/"
    if cfg.Next != nil {
        fullPath = *cfg.Next
    }

    values, err := getLocationArea(&fullPath, cfg)
    if err != nil {
        return fmt.Errorf("Could not find areas: %w", err)
    }

    for i, val := range values.Results {
        fmt.Printf("%d. %s -> %s\n", i, val.Name, val.Url)
    }
	return nil
}

func mapbCallback(cfg *config) error {
	fullPath := baseURL + "location-area/"
    if cfg.Previous != nil {
        fullPath = *cfg.Previous
    }

    values, err := getLocationArea(&fullPath, cfg)
    if err != nil {
        return fmt.Errorf("Could not find areas: %w", err)
    }

    for i, val := range values.Results {
        fmt.Printf("%d. %s -> %s\n", i, val.Name, val.Url)
    }
	return nil
}
