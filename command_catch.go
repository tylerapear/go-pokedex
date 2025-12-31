package main

import (
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	
	if len(args) < 1 {
		return fmt.Errorf("usage: catch <pokemon_name>")
	}

	pokemonNameStr := args[0]
	pokemonName := &pokemonNameStr
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonNameStr)

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	PercentChanceToCatch := float64(1) - (float64(pokemonResp.Base_Experience) / 350)
	PercentChanceToCatch = math.Round(PercentChanceToCatch * 100) / 100

	randomNum := rand.Float64()

	if randomNum < PercentChanceToCatch {
		fmt.Printf("Congratulations! You caught a %s! This was added to your Pokedex.\n", pokemonResp.Name,)
		cfg.pokedex[pokemonResp.Name] = Pokemon{
			Name: pokemonResp.Name,
			Count: cfg.pokedex[pokemonResp.Name].Count + 1,
		}
	} else {
		fmt.Printf("Oh no! The %s broke free!\n", pokemonResp.Name)
	}

	return nil
}