package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	lowerCaseString := strings.ToLower(text)
	trimmedString := strings.Trim(lowerCaseString, " ")
	splitString := strings.Split(trimmedString, " ")
	return splitString
}

func main() {
	fmt.Println(cleanInput("Hello, World!"))
}
