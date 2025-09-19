package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Cadimodev/pokedexcli/internal/pokecache"
)

//https://pokeapi.co/api/v2/pokemon/{id or name}/

func (c *Client) GetPokemonData(pokemonName string, cache *pokecache.Cache) (Pokemon, error) {

	url := baseURL + "/pokemon/" + pokemonName

	//TODO: delete this line
	fmt.Println("URL -> ", url)

	data, ok := cache.Get(url)
	if !ok {
		fmt.Println("Cache not found, making request...")

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		resp, err := c.httpCLient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, err
		}

	} else {
		fmt.Println("Reading from cache...")
	}

	cache.Add(url, data)

	pokemonData := Pokemon{}
	err := json.Unmarshal(data, &pokemonData)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonData, nil
}
