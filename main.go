package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("PokéCli > ")
		scanner.Scan()
        words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display help message",
			callback:    helpCallback,
		},
		"exit": {
			name:        "exit",
			description: "Exit the PokéCli",
			callback:    exitCallback,
		},
	}
}

func helpCallback() error {
	fmt.Printf("== Help menu ==\n\n")
	fmt.Printf("Command\t\tDescription\n------------------------------\n")
	for k, v := range getCommands() {
		fmt.Printf("%s\t\t%s\n", k, v.description)
	}
	fmt.Println("------------------------------")
	return nil
}

func exitCallback() error {
	os.Exit(0)
	return nil
}
