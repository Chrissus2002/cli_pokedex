package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Chrissus2002/cli_pokedex/internal/pokecache"
)

func PokeMap(url string, cache pokecache.Cache) LocationsAreas{
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()


	if val, ok := cache.Get(url); ok{
		areas := LocationsAreas{}
		err := json.Unmarshal(val, &areas)
		if err != nil{
			fmt.Println(err)
		}
		return areas
	}

	var areas LocationsAreas
	err = json.Unmarshal(body, &areas)
	if err != nil {
		fmt.Println(err)
	}
	
	cache.Add(url, body)
	return areas
}