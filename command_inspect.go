package main

import (
	"fmt"
)

func commandInspect(p *config, args ...string) error{
	if len(args) < 1{
		fmt.Println("Please enter a pokemon!")
		return nil
	}
	
	pokemon_name := args[0]
	if el, ok := p.caughtPokemon[pokemon_name]; ok{
		fmt.Println("Name:", pokemon_name)
		fmt.Println("Height:", el.Height)
		fmt.Println("Weight:", el.Weight)
		fmt.Println("Stats:")
		for _, stat := range el.Stats{
			fmt.Printf("  -%s: %d\n",stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, types := range el.Types{
			fmt.Println(" -", types.Type.Name)
		}
		return nil
	}else{
		fmt.Println("You have not caught this pokemon yet.")
	}
	return nil
}