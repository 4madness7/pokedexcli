package main

import "fmt"

func helpCallback(cfg *config, args ...string) error {
    if len(args) > 0 {
        return fmt.Errorf("Error: too many arguments. 0 arguments expected, %d provided.", len(args))
    }

	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("== Help menu ==\n\n")
	fmt.Printf("Command\t\tDescription\n------------------------------\n")
	for k, v := range getCommands() {
		fmt.Printf("%s\t\t%s\n", k, v.description)
	}
	fmt.Println("------------------------------")
	return nil
}

