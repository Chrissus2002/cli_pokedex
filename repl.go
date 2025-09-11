package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Chrissus2002/cli_pokedex/internal/pokecache"
)

type cliCommand struct{
	name string
	description string
	callback func(*config, ...string) error
}

type config struct{
	Cache pokecache.Cache
	Next *string `json:"next"`
	Previous *string `json:"previous"`
}

func startREPL(){
	scanner := bufio.NewScanner(os.Stdin)
	conf := config{
		Cache: *pokecache.NewCache(time.Minute * 5),
		Next: nil,
		Previous: nil,
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		cleaned_text := cleanInput(scanner.Text())
		if len(cleaned_text) == 0{
			continue
		}
		command_words := cleaned_text[0]
		args := []string{}
		if len(cleaned_text) > 1{
			args = cleaned_text[1:]
		}


		command, ok := getCommands()[command_words]
		if ok{
			err := command.callback(&conf,args...)
			if err != nil{
				fmt.Println(err)
			}
			continue
		} else{
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowerCaseString := strings.ToLower(text)
	trimmedString := strings.Trim(lowerCaseString, " ")
	splitString := strings.Split(trimmedString, " ")
	return splitString
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"exit":{
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help":{
			name: "help",
			description: "Display a help message",
			callback: commandHelp,
		},
		"map":{
			name: "map",
			description: "Display the next 20 areas in the Pokedex",
			callback: commandMap,
		},
		"mapb":{
			name: "mapb",
			description: "Display the previous 20 areas in the Pokedex",
			callback: commandMapPrev,
		},
		"explore":{
			name: "explore",
			description: "Explore a location area that you see in the Pokedex map",
			callback: commandExplore,
		},
	}
}