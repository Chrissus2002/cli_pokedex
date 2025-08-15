package main

import (
	"fmt"
	"github.com/Chrissus2002/cli_pokedex/internal/pokeapi"
)

func commandMap(p *config) error {
	url := pokeapi.Base_url + "/location-area/"
	if p.Next == nil{
		p.Next = &url
	}
	locs := pokeapi.PokeMap(*p.Next)
	for _, el := range locs.Results{
		fmt.Println(el.Name)
	}
	p.Next = &locs.Next
	p.Previous = locs.Previous
	return nil
}

func commandMapPrev(p *config) error {
	if p.Previous == nil{
		fmt.Println("you're on the first page")
	}else{
		locs := pokeapi.PokeMap(*p.Previous)
		for _,el := range locs.Results{
			fmt.Println(el.Name)
		}
		p.Next = &locs.Next
		p.Previous = locs.Previous
	}
	return nil
}