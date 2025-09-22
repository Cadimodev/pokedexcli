package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("invalid param")

	}

	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("%s is not caught. no info available")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Printf("-%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("-%s\n", typeInfo.Type.Name)
	}

	return nil
}
