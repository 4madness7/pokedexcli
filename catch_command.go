package main

import (
	"fmt"
	"math/rand"
)

func catchCallback(cfg *config, args ...string) error {
	if len(args) > 1 {
		return fmt.Errorf("Error: too many arguments. 1 arguments expected, %d provided.", len(args))
	}
	if len(args) < 1 {
		return fmt.Errorf("Error: too few arguments. 1 arguments expected, %d provided.", len(args))
	}

	pokemonName := args[0]

	pokemon, err := cfg.Client.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	catchRate := -.22*float32(pokemon.BaseExperience) + 110

    n := rand.Float32()*100

    // if n > catchRate, pokemon is not caught
    if n > catchRate {
        fmt.Printf("%s escaped!\n", pokemon.Name)
        return nil
    }

    fmt.Printf("%s was caught!\n", pokemon.Name)
    cfg.CaughtPokemons[pokemon.Name] = pokemon
	return nil
}
