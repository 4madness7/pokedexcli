package main

import "fmt"

func inspectCallback(cfg *config, args ...string) error {
	if len(args) > 1 {
		return fmt.Errorf("Error: too many arguments. 1 arguments expected, %d provided.", len(args))
	}
	if len(args) < 1 {
		return fmt.Errorf("Error: too few arguments. 1 arguments expected, %d provided.", len(args))
	}

	pokemonName := args[0]

	pokemon, ok := cfg.CaughtPokemons[pokemonName]
	if !ok {
		return fmt.Errorf("Pokémon not caught!")
	}

    fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n",
		pokemon.Name,
		pokemon.Height,
		pokemon.Weight,
	)

    fmt.Println("Stats:")
    for _, s := range pokemon.Stats {
        fmt.Printf("    - %s: %d\n", s.Stat.Name, s.BaseStat)
    }

    fmt.Println("Type(s):")
    for _, t := range pokemon.Types {
        fmt.Printf("    - %s\n", t.Type.Name)
    }
	return nil
}
