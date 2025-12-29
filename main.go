package main

import (
	"fmt"
	"os"
	"bufio"
	"time"

	"github.com/tylerapear/go-pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient 			pokeapi.Client
	nextLocationsURL 		*string
	previousLocationsURL 	*string
}

type cliCommand struct {
	name		string
	description	string
	callback 	func(*config) error
	config 		*config
}


var commands map[string]cliCommand

func main() {

	config := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
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
	}
	
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedText := cleanInput(scanner.Text())
		if len(cleanedText) == 0 {
			continue
		}

		commandName := cleanedText[0]

		command, exists := commands[commandName]
		if exists {
			err := command.callback(config)
			if err != nil {
				fmt.Println("Error executing command:", err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}

		
	}

}
