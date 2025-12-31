package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {

	if len(args) < 1 {
		return fmt.Errorf("usage: inspect <pokemon_name>")
	}

	pokemonNameStr := args[0]
	pokemonName := &pokemonNameStr

	_, exists := cfg.pokedex[pokemonNameStr]
	if !exists {
		fmt.Printf("You have not caught a %s yet.\n", pokemonNameStr)
		return nil
	}

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", pokemonResp.Name)
	fmt.Printf("Height: %s\n", string(pokemonResp.Height))
	fmt.Printf("Weight: %s\n", pokemonResp.Weight)

	fmt.Printf("Stats:\n")
	for _, stat := range pokemonResp.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.Base_Stat)
	}	

	fmt.Printf("Types:\n")
	for _, t := range pokemonResp.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}	

	return nil

}