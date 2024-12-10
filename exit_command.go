package main

import (
    "fmt"
    "os"
)

func exitCallback(cfg *config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
