package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type config struct {
	Client   *http.Client
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func main() {
	cfg := &config{
		Client: &http.Client{},
	}
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
			err := command.callback(cfg)
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
		"map": {
			name:        "map",
			description: "Move to next 20 iteration of location",
			callback: mapfCallback,
		},
		"mapb": {
			name:        "mapb",
			description: "Move to prev 20 iteration of location",
			callback: mapbCallback,
		},
	}
}

