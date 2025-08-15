package main

import(
	"fmt"
	"strings"
	"os"
	"bufio"
)

type cliCommand struct{
	name string
	description string
	callback func(*config) error
}

type config struct{
	Next *string `json:"next"`
	Previous *string `json:"previous"`
}

func startREPL(){
	scanner := bufio.NewScanner(os.Stdin)
	conf := config{
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


		command, ok := getCommands()[command_words]
		if ok{
			err := command.callback(&conf)
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
	}
}