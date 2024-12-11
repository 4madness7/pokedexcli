package pokeapi

import (
	"net/http"
	"time"

	"github.com/4madness7/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	Cache      pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		Cache: pokecache.NewCache(timeout),
	}
}
