package main

import (
	"fmt"
)

func commandHelp(cfg *httpConfig) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, value := range getCommands() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}

	return nil
}
