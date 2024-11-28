package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex")
	fmt.Println("Usage:")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandExit() error {
	fmt.Println("Exiting the Pokedex.")
	return nil
}

func main() {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex> ")
		if scanner.Scan() {
			input := scanner.Text()

			if cmd, exists := commands[input]; exists {
				if err := cmd.callback(); err != nil {
					fmt.Println("Error executing command:", err)
				}

				switch input {
				case "help":

				case "exit":
					return
				}
			}
		}
	}

}
