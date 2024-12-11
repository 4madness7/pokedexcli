package main

import (
	"encoding/json"
	"fmt"
	"github.com/4madness7/pokedexcli/internal/pokeapi"
)

func mapfCallback(cfg *config) error {
	var locations *pokeapi.DataLocationArea

	if cfg.Next != nil {
		data, exists := cfg.Cache.Get(*cfg.Next)
		if exists {
			locationsNew := pokeapi.DataLocationArea{}
			err := json.Unmarshal(data, &locationsNew)
			if err != nil {
				return fmt.Errorf("Could not find areas: %w", err)
			}
			locations = &locationsNew
		}
	}

	if locations == nil {
		fmt.Println("=== Making a request to Pokeapi! ===")
		locationsResp, err := cfg.Client.ListLocations(cfg.Next, &cfg.Cache)
		if err != nil {
			return fmt.Errorf("Could not find areas: %w", err)
		}
		locations = &locationsResp
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

    printingLocations(locations)

	return nil
}

func mapbCallback(cfg *config) error {
	if cfg.Previous == nil {
		return fmt.Errorf("Error: you are in the first page")
	}

	var locations *pokeapi.DataLocationArea

	data, exists := cfg.Cache.Get(*cfg.Previous)
	if exists {
		locationsNew := pokeapi.DataLocationArea{}
		err := json.Unmarshal(data, &locationsNew)
		if err != nil {
			return fmt.Errorf("Could not find areas: %w", err)
		}
		locations = &locationsNew
	}
	if locations == nil {
		fmt.Println("=== Making a request to Pokeapi! ===")
		locationsResp, err := cfg.Client.ListLocations(cfg.Previous, &cfg.Cache)
		if err != nil {
			return fmt.Errorf("Could not find areas: %w", err)
		}
		locations = &locationsResp
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

    printingLocations(locations)

	return nil
}

func printingLocations(data *pokeapi.DataLocationArea) {
	for _, val := range data.Results {
		fmt.Println(val.Name)
	}
}
