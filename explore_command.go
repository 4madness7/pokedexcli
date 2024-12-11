package main

import (
	"encoding/json"
	"fmt"

	"github.com/4madness7/pokedexcli/internal/pokeapi"
)

func exploreCallback(cfg *config, args []string) error {
    if len(args) > 1 {
        return fmt.Errorf("Error: too many arguments. 1 arguments expected, %d provided.", len(args))
    }
    if len(args) < 1 {
        return fmt.Errorf("Error: too few arguments. 1 arguments expected, %d provided.", len(args))
    }

    locationToExplore := args[0]

	var location *pokeapi.ExploredLocationArea

    fmt.Printf("Exploring %s...\n", locationToExplore)

	data, exists := cfg.Cache.Get(locationToExplore)
	if exists {
		locationNew := pokeapi.ExploredLocationArea{}
		err := json.Unmarshal(data, &locationNew)
		if err != nil {
			return fmt.Errorf("Could not find area '%s': %w", locationToExplore, err)
		}
		location = &locationNew
	}
	if location == nil {
		locationResp, err := cfg.Client.ExploreLocationArea(locationToExplore, &cfg.Cache)
		if err != nil {
			return fmt.Errorf("Could not find area '%s': %w", locationToExplore, err)
		}
		location = &locationResp
	}

    fmt.Printf("Found Pokémon:\n")
    for _, pokemon := range location.PokemonEncounters {
        fmt.Printf("- %s\n", pokemon.Pokemon.Name)
    }

    return nil
}
