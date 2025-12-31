package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

func (c *Client) GetPokemon(pokemon *string) (RespShallowPokemon, error) {

	url := baseURL + fmt.Sprintf("/pokemon/%s", *pokemon)

	// First, check cache. If not in cache, make request and add to cache.
	dat := []byte{}
	cacheEntry, exists := c.pokecache.Get(url)
	if exists {
		dat = cacheEntry
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowPokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowPokemon{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowPokemon{}, err
		}

		c.pokecache.Add(url, dat)
	}

	// Then, unmarshal the data and return it.
	pokemonResp := RespShallowPokemon{}
	err := json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	return pokemonResp, nil

}