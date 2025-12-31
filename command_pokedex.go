package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {

	if len(cfg.pokedex) == 0 {
		fmt.Println("You have not caught any Pokemon yet.")
		return nil
	}

	fmt.Println("Pokemon in your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %s (%d)\n", pokemon.Name, pokemon.Count)
	}

	return nil

}