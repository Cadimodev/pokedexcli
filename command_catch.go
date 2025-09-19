package main

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	catchCoef = 40 //Used for calculate if pokemon catch was successful
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("invalid param")

	}
	pokemonName := args[0]
	pokemonData, err := cfg.pokeapiClient.GetPokemonData(pokemonName, cfg.pokeCache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonData.Name)

	res := rand.Intn(pokemonData.BaseExperience)

	if res > catchCoef {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonData.Name)
	cfg.caughtPokemon[pokemonData.Name] = pokemonData

	return nil
}
