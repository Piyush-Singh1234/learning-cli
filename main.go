package main

import (
	"bufio"
	// "encoding/json"
	"fmt"
	// "io"
	// "log"
	// "net/http"
	"os"
	// "strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}


var climap map[string]cliCommand

func commandExit(cfg *config) error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


func commandHelp(cfg *config) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _,cmd := range climap{
		fmt.Printf("%v : %v\n",cmd.name,cmd.description)
	}
	return nil
}

var apicall = "https://pokeapi.co/api/v2/location-area/"

func commandmap(cfg *config) error {
	var nxt *string
	var prev *string
	var Location[]string
	var Error error 
	if cfg.nextURL == nil{
		nxt,prev,Location,Error = fetchLocationAreas(&apicall)
	}else{
		nxt,prev,Location,Error = fetchLocationAreas(cfg.nextURL)
	}
	if Error!=nil{
		return Error
	}
	for _,loc := range(Location){
		fmt.Println(loc)
	}
	cfg.nextURL=nxt
	cfg.prevURL=prev
	return nil
}

func commandmapb(cfg *config) error {
	var nxt *string
	var prev *string
	var Location[]string
	var Error error 
	if cfg.prevURL == nil{
		fmt.Println("you're on the first page")
		return nil
	}else{
		nxt,prev,Location,Error = fetchLocationAreas(cfg.prevURL)
	}
	if Error!=nil{
		return Error
	}
	for _,loc := range(Location){
		fmt.Println(loc)
	}
	cfg.nextURL=nxt
	cfg.prevURL=prev
	return nil
}

func main() {
	
	var cfg = &config{}
	cfg.prevURL = nil
	cfg.nextURL = nil

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
		"map":{
			name: "map",
			description: "map command",
			callback: commandmap,
		},
		"mapb":{
			name: "mapb",
			description: "mapb command",
			callback: commandmapb,
		},
	}




	for {
		fmt.Print("pokedex > ")

		if !scanner.Scan() {
			break
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmd, ok := climap[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.callback(cfg); err != nil {
			fmt.Println("Error:", err)
		}
	}

}
