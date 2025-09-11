package main

import (
	"fmt"
	"math/rand"

	"github.com/Chrissus2002/cli_pokedex/internal/pokemon"
)

func commandCatch(p *config, args ...string) error{
	pokemon_name := args[0]
	pokemonInfo, err := pokemon.CatchPokemon(pokemon_name)
	if err != nil{
		fmt.Println(err)
	}

	catchChance := rand.Intn(pokemonInfo.BaseExperience)

	fmt.Println("Throwing a Pokeball at", pokemon_name + "...")
	if catchChance > 50{
		fmt.Println(pokemon_name, "escaped!")
		return nil
	}
	fmt.Println(pokemon_name, "was caught!")

	p.caughtPokemon[pokemon_name] = pokemonInfo
	return nil
}