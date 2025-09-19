package main

import (
	"time"

	"github.com/Cadimodev/pokedexcli/internal/pokeapi"
	"github.com/Cadimodev/pokedexcli/internal/pokecache"
)

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second)
	cache := pokecache.NewCache(6 * time.Hour)

	cfg := &config{
		pokeapiClient: pokeClient,
		pokeCache:     cache,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}
