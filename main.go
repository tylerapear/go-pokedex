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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for command := range commands {
		fmt.Printf("%s: %s\n", commands[command].name, commands[command].description)
	}
	return nil
}