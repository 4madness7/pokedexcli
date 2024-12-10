package main

import (
    "fmt"
)

func mapfCallback(cfg *config) error {
    locations, err := cfg.Client.ListLocations(cfg.Next)
    if err != nil {
        return fmt.Errorf("Could not find areas: %w", err)
    }

    cfg.Next = locations.Next
    cfg.Previous = locations.Previous

    for i, val := range locations.Results {
        fmt.Printf("%d. %s -> %s\n", i, val.Name, val.Url)
    }
	return nil
}

func mapbCallback(cfg *config) error {
    if cfg.Previous == nil {
        return fmt.Errorf("Error: you are in the first page")
    }

    locations, err := cfg.Client.ListLocations(cfg.Previous)
    if err != nil {
        return fmt.Errorf("Could not find areas: %w", err)
    }

    cfg.Next = locations.Next
    cfg.Previous = locations.Previous

    for i, val := range locations.Results {
        fmt.Printf("%d. %s -> %s\n", i, val.Name, val.Url)
    }
	return nil
}
