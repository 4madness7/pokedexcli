package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	baseURL = "https://pokeapi.co/api/v2/"
)

type DataLocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

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

func getLocationArea(path *string, cfg *config) (DataLocationArea, error) {
	res, err := cfg.Client.Get(*path)
	if err != nil {
		return DataLocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return DataLocationArea{}, err
	}

	var values DataLocationArea
	err = json.Unmarshal(data, &values)
	if err != nil {
		return DataLocationArea{}, err
	}

    cfg.Next = values.Next
    cfg.Previous = values.Previous

	return values, nil
}

func mapfCallback(cfg *config) error {
	fullPath := baseURL + "location-area/"
    if cfg.Next != nil {
        fullPath = *cfg.Next
    }

    values, err := getLocationArea(&fullPath, cfg)
    if err != nil {
        return fmt.Errorf("Could not find areas: %w", err)
    }

    for i, val := range values.Results {
        fmt.Printf("%d. %s -> %s\n", i, val.Name, val.Url)
    }
	return nil
}

func mapbCallback(cfg *config) error {
	fullPath := baseURL + "location-area/"
    if cfg.Previous != nil {
        fullPath = *cfg.Previous
    }

    values, err := getLocationArea(&fullPath, cfg)
    if err != nil {
        return fmt.Errorf("Could not find areas: %w", err)
    }

    for i, val := range values.Results {
        fmt.Printf("%d. %s -> %s\n", i, val.Name, val.Url)
    }
	return nil
}

func helpCallback(cfg *config) error {
	fmt.Printf("== Help menu ==\n\n")
	fmt.Printf("Command\t\tDescription\n------------------------------\n")
	for k, v := range getCommands() {
		fmt.Printf("%s\t\t%s\n", k, v.description)
	}
	fmt.Println("------------------------------")
	return nil
}

func exitCallback(cfg *config) error {
	os.Exit(0)
	return nil
}
