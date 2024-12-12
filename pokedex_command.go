package main

import "fmt"

func pokedexCallback(cfg *config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("Error: too many arguments. 0 arguments expected, %d provided.", len(args))
	}

	if len(cfg.CaughtPokemons) == 0 {
		return fmt.Errorf("No Pokémons caught yet.")
	}

	fmt.Println("Your Pokédex:")
	for k := range cfg.CaughtPokemons {
		fmt.Printf("    - %s\n", k)
	}

	return nil
}
