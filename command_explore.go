package main

import (
	"fmt"
	//"errors"
)


func commandExplore(cfg *config, args []string) (error) {

	if len(args) < 1 {
		return fmt.Errorf("usage: explore <location_name>")
	}

	locationName := args[0]

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon:")

	urlStr := fmt.Sprintf(locationName)
	url := &urlStr

	locationResp, err := cfg.pokeapiClient.GetLocation(url)
	if err != nil {
		return err
	}

	for _, pokemon := range locationResp.Pokemon_encounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	
	return nil
}