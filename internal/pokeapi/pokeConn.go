package pokeapi

import(
	"fmt"
	"encoding/json"
	"net/http"
	"io"
)

func PokeMap(url string) LocationsAreas{
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	var areas LocationsAreas
	err = json.Unmarshal(body, &areas)
	if err != nil {
		fmt.Println(err)
	}
	return areas
}