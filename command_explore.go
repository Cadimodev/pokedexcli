package main

import (
	"fmt"
)

func commandExplore(cfg *httpConfig) error {

	fmt.Printf("Exploring %s ...\n", cfg.param)

	pokemonNames, err2 := cfg.pokeapiClient.GetLocationData(cfg.param, cfg.pokeCache)
	if err2 != nil {
		return err2
	}

	fmt.Println("Found Pokemon:")
	for _, pokemonName := range pokemonNames {
		fmt.Println("- ", pokemonName)
	}

	return nil
}
