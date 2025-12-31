package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

func (c *Client) GetLocation(location *string) (RespLocation, error) {

	url := baseURL + fmt.Sprintf("/location-area/%s", *location)

	// First, check cache. If not in cache, make request and add to cache.
	dat := []byte{}
	cacheEntry, exists := c.pokecache.Get(url)
	if exists {
		dat = cacheEntry
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocation{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocation{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespLocation{}, err
		}

		c.pokecache.Add(url, dat)
	}

	// Then, unmarshal the data and return it.
	locationResp := RespLocation{}
	err := json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespLocation{}, err
	}

	return locationResp, nil

}