package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/4madness7/pokedexcli/internal/pokeapi"
)

type config struct {
	Client         pokeapi.Client
	CaughtPokemons map[string]pokeapi.Pokemon
	Next           *string
	Previous       *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func main() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		CaughtPokemons: map[string]pokeapi.Pokemon{},
		Client:         client,
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
			err := command.callback(cfg, words[1:]...)
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
			callback:    mapfCallback,
		},
		"mapb": {
			name:        "mapb",
			description: "Move to prev 20 iteration of location",
			callback:    mapbCallback,
		},
		"explore": {
			name:        "explore",
			description: "Explore specified area. (eg. explore <area-name>)",
			callback:    exploreCallback,
		},
		"catch": {
			name:        "catch",
			description: "Try catching specified Pokémon (eg. catch <pokémon-name>)",
			callback:    catchCallback,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect specified Pokémon if caught. (eg. inspect <pokémon-name>)",
			callback:    inspectCallback,
		},
	}
}
