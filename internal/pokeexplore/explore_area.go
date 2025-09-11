package pokeexplore

import(
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"github.com/Chrissus2002/cli_pokedex/internal/pokeapi"
	"github.com/Chrissus2002/cli_pokedex/internal/pokecache"
)

func PokeExplore(area string, cache pokecache.Cache) AreaDetailedResp{
	url := pokeapi.Base_url + "/location-area/" + area
	res, err := http.Get(url)
	if err != nil{
		fmt.Println("Error with request:", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil{
		fmt.Println("Error reading response body:", err)
	}
	defer res.Body.Close()

	if val, ok := cache.Get(url); ok{
		AreaInfo := AreaDetailedResp{}
		err := json.Unmarshal(val, &AreaInfo)
		if err != nil{
			fmt.Println(err)
		}
		return AreaInfo
	}

	AreaInfo := AreaDetailedResp{}
	err = json.Unmarshal(body, &AreaInfo)
	if err != nil{
		fmt.Println("Error with json body structure: ", err)
	}

	cache.Add(url, body)
	return AreaInfo
}
