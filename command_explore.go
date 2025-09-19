package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	locationName := args[0]

	fmt.Printf("Exploring %s ...\n", locationName)

	pokemonNames, err2 := cfg.pokeapiClient.GetLocationData(locationName, cfg.pokeCache)
	if err2 != nil {
		return err2
	}

	fmt.Println("Found Pokemon:")
	for _, pokemonName := range pokemonNames {
		fmt.Println("- ", pokemonName)
	}

	return nil
}
