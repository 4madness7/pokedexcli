package main

import (
    "fmt"
    "os"
)

func exitCallback(cfg *config, args ...string) error {
    if len(args) > 0 {
        return fmt.Errorf("Error: too many arguments. 0 arguments expected, %d provided.", len(args))
    }

	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
