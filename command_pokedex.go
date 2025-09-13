package main

import "fmt"

func commandPokedex(p *config, args ...string) error{
	for key, _ := range p.caughtPokemon{
		fmt.Println(" -", key)
	}
	return nil
}