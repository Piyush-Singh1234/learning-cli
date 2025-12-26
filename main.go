package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct{
	name string
	description string
	callback func() error
}

var climap map[string]cliCommand

func commandExit() error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandHelp() error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _,cmd := range climap{
		fmt.Printf("%v : %v\n",cmd.name,cmd.description)
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	climap =  map[string]cliCommand{
		"help":{
			name : "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit":{
			name: "exit",
			description: "Exits the terminal",
			callback: commandExit,
		},
	}

	for {
		fmt.Print("pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		cmd, ok := climap[input]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.callback(); err != nil {
			fmt.Println("Error:", err)
		}
	}

}
