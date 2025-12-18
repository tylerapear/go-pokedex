package main

import (
	"fmt"
	"os"
	"bufio"
)

type config struct {
	Next 		string
	Previous 	string
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
		Next: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		Previous: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	}

	commands = map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
			config: config,
		},
		"help": {
			name: "help",
			description: "Display help message",
			callback: commandHelp,
			config: config,
		},
		"map": {
			name: "map",
			description: "Display next page of map information",
			callback: commandMapf,
			config: config,
		},
		"mapb": {
			name: "mapb",
			description: "Display previous page of map information",
			callback: commandMapb,
			config: config,
		},
	}
	
	scanner := bufio.NewScanner(os.Stdin)

	for ; ; {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedText := cleanInput(scanner.Text())

		command, ok := commands[cleanedText[0]]
		if !ok {
			fmt.Println("Unknown command")
		}

		err := command.callback(command.config)
		if err != nil {
			fmt.Println("Error executing command:", err)
		}
	}

}
