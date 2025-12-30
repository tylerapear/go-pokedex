package pokeapi

import (
	"net/http"
	"time"

	"github.com/tylerapear/go-pokedex/internal/pokecache"
)

// Client
type Client struct {
	httpClient 	http.Client
	pokecache 	pokecache.Cache
}

// NewClient
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokecache: pokecache.NewCache(10 * time.Minute),
	}
}

func (c Client) GetCache() pokecache.Cache {
	return c.pokecache
}