package main

import (
	"fmt"
)

func commandPCache(cfg *config, args []string) (error) {
	fmt.Println("\nPokeCache Entries:")
	fmt.Println("")

	entries := cfg.pokeapiClient.GetCache().GetAll()
	for key, entry := range entries {
		fmt.Printf("Key: %s\nValue: %s\nCreated At: %s\n", key, string(entry.GetVal()), entry.GetCreatedAt().String())
		fmt.Println("")
	}
	return nil
}