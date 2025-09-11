package main

import(
	"fmt"

	"github.com/Chrissus2002/cli_pokedex/internal/pokeexplore"
)

func commandExplore(p *config, args ...string) error{
	if len(args) != 1{
		fmt.Println("Enter just a single area")
		return nil
	}

	area_name := args[0]
	fmt.Println("Exploring", area_name, "...")
	fmt.Println("Found pokemon:")
	PokemonEncounters := pokeexplore.PokeExplore(area_name, p.Cache)

	for _, el := range PokemonEncounters.PokemonEncounters{
		fmt.Println("-", el.Pokemon.Name)
	}

	return nil
}