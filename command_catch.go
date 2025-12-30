package main

import (
	"fmt"
	"encoding/json"

	"github.com/tylerapear/go-pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, args []string) error {
	
	if len(args) < 1 {
		return fmt.Errorf("usage: catch <pokemon_name>")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	dat, err := httpJSONGet(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName))
	if err != nil {
		return err
	}

	pokemonResp := pokeapi.RespShallowPokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return err
	}

	fmt.Printf("Congratulations! You caught a %s, with base experience: %d\n", pokemonResp.Name, pokemonResp.Base_Experience)

	return nil
}