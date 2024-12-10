package main

import "fmt"

func helpCallback(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("== Help menu ==\n\n")
	fmt.Printf("Command\t\tDescription\n------------------------------\n")
	for k, v := range getCommands() {
		fmt.Printf("%s\t\t%s\n", k, v.description)
	}
	fmt.Println("------------------------------")
	return nil
}

