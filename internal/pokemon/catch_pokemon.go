package pokemon

import(
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	
	"github.com/Chrissus2002/cli_pokedex/internal/pokeapi"
)

func CatchPokemon(pokemon_name string) (PokemonData, error){
	url := pokeapi.Base_url + "/pokemon/" + pokemon_name

	res, err := http.Get(url)
	if err != nil{
		fmt.Println(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil{
		fmt.Println(err)
	}

	var pokemonStats PokemonData
	err = json.Unmarshal(body, &pokemonStats)
	if err != nil{
		fmt.Println(err)
	}

	return pokemonStats,nil
}