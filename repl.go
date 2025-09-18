package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Cadimodev/pokedexcli/internal/pokeapi"
	"github.com/Cadimodev/pokedexcli/internal/pokecache"
)

type httpConfig struct {
	pokeapiClient    pokeapi.Client
	pokeCache        *pokecache.Cache
	param            string
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *httpConfig) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		inputRes := cleanInput(input)

		if len(inputRes) == 0 {
			fmt.Println("Error: invalid input!")
			continue
		}

		commandName := inputRes[0]
		if len(inputRes) >= 2 {
			cfg.param = inputRes[1]
		}

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

	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*httpConfig) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
