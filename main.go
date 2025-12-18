package main

import (
	"fmt"
	"os"
	"bufio"
)

type cliCommand struct {
	name		string
	description	string
	callback 	func() error
}

var commands map[string]cliCommand

func main() {

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

		command.callback()
	}

}
