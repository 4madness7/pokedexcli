package main

import (
	"fmt"
	"github.com/4madness7/pokedexcli/internal/pokeapi"
)

func mapfCallback(cfg *config, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("Error: too many arguments. 0 arguments expected, %d provided.", len(args))
	}

	locations, err := cfg.Client.ListLocations(cfg.Next)
	if err != nil {
		return fmt.Errorf("Could not find areas: %w", err)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	printingLocations(&locations)

	return nil
}

func mapbCallback(cfg *config, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("Error: too many arguments. 0 arguments expected, %d provided.", len(args))
	}

	if cfg.Previous == nil {
		return fmt.Errorf("Error: you are in the first page")
	}

	locations, err := cfg.Client.ListLocations(cfg.Previous)
	if err != nil {
		return fmt.Errorf("Could not find areas: %w", err)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	printingLocations(&locations)

	return nil
}

func printingLocations(data *pokeapi.DataLocationArea) {
	for _, val := range data.Results {
		fmt.Println(val.Name)
	}
}
