package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// List Locations
func (c *Client) ListLocations(pageURL *string) (RespShallowLocaions, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// First, check cache. If not in cache, make request and add to cache.
	dat := []byte{}
	cacheEntry, exists := c.pokecache.Get(url)
	if exists {
		dat = cacheEntry
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocaions{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocaions{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocaions{}, err
		}

		c.pokecache.Add(url, dat)
	}

	// Then, unmarshal the data and return it.
	locationsResp := RespShallowLocaions{}
	err := json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocaions{}, err
	}

	return locationsResp, nil
}