package main

import (
	"time"

	"github.com/tylerapear/go-pokedex/internal/pokeapi"
)

type Pokemon struct {
	Name string
	Count int
}

func main() {

	config := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
		pokedex: make(map[string]Pokemon),
	}
	startRepl(config)

}
