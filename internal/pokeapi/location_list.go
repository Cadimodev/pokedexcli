package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Cadimodev/pokedexcli/internal/pokecache"
)

func (c *Client) ListLocations(pageURL *string, cache *pokecache.Cache) (RespShallowLocations, error) {

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	fmt.Println("URL -> ", url)
	data, ok := cache.Get(url)
	if !ok {

		fmt.Println("Cache not found, making request...")
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, err
		}

		resp, err := c.httpCLient.Do(req)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}

	} else {
		fmt.Println("Reading from cache...")
	}

	locationsResp := RespShallowLocations{}
	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	cache.Add(url, data)

	return locationsResp, nil
}

func (c *Client) GetLocationData(locationName string, cache *pokecache.Cache) ([]string, error) {
	url := baseURL + "/location-area/" + locationName

	fmt.Println("URL -> ", url)

	data, ok := cache.Get(url)
	if !ok {

		fmt.Println("Cache not found, making request...")
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return []string{}, err
		}

		resp, err := c.httpCLient.Do(req)
		if err != nil {
			return []string{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return []string{}, err
		}

	} else {
		fmt.Println("Reading from cache...")
	}

	cache.Add(url, data)

	locationResp := RespLocationData{}
	err := json.Unmarshal(data, &locationResp)
	if err != nil {
		return []string{}, err
	}

	pokemonNames := make([]string, 0, len(locationResp.PokemonEncounters))
	for _, encounter := range locationResp.PokemonEncounters {
		fmt.Println("pokemon : ", encounter.Pokemon.Name)
		pokemonNames = append(pokemonNames, encounter.Pokemon.Name)
	}

	return pokemonNames, nil
}
