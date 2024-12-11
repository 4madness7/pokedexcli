package main

import (
	"fmt"
)

func exploreCallback(cfg *config, args ...string) error {
	if len(args) > 1 {
		return fmt.Errorf("Error: too many arguments. 1 arguments expected, %d provided.", len(args))
	}
	if len(args) < 1 {
		return fmt.Errorf("Error: too few arguments. 1 arguments expected, %d provided.", len(args))
	}

	locationToExplore := args[0]

	fmt.Printf("Exploring %s...\n", locationToExplore)

	location, err := cfg.Client.ExploreLocationArea(locationToExplore)
	if err != nil {
		return fmt.Errorf("Could not find area '%s': %w", locationToExplore, err)
	}

	fmt.Printf("Found Pokémon:\n")
	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
