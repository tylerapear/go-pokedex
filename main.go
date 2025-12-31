package main

import (
	"fmt"
	"os"
	"bufio"
	"time"

	"github.com/tylerapear/go-pokedex/internal/pokeapi"
)

type Pokemon struct {
	Name string
	Count int
}

type config struct {
	pokeapiClient 			pokeapi.Client
	nextLocationsURL 		*string
	previousLocationsURL 	*string
	pokedex 				map[string]Pokemon
}

type cliCommand struct {
	name		string
	description	string
	callback 	func(*config, []string) error
	config 		*config
}

var commands map[string]cliCommand

func main() {

	config := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
		pokedex: make(map[string]Pokemon),
	}

	commands = map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Display help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Display next page of map information",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Display previous page of map information",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Explore a location and list available Pokemon. Usage: explore <location_name>",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Attempt to catch a Pokemon. Usage: catch <pokemon_name>",
			callback: commandCatch,
		},
		"pcache": {
			name: "pcache",
			description: "Print out the current PokeCache entries",
			callback: commandPCache,
		},
	}
	
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedText := cleanInput(scanner.Text())
		if len(cleanedText) == 0 {
			continue
		}

		args := []string{}
		commandName := cleanedText[0]
		if len(cleanedText) > 1 {
			args = cleanedText[1:]
		}

		command, exists := commands[commandName]
		if exists {
			startTime := time.Now()
			err := command.callback(config, args)
			elapsed := time.Since(startTime)
			if contains(args, "-v") || contains(args, "--verbose") {
				fmt.Printf("Command '%s' executed in %v microseconds\n", commandName, elapsed.Microseconds())
			}
			if err != nil {
				fmt.Println("Error executing command:", err)
			}
		} else {
			fmt.Printf("'%s': Unknown command\n", commandName)
			continue
		}

		
	}

}
