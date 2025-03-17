package pokeapi

import (
	"net/http"
	"time"

	"github.com/janisgee/pokedexcli-app/internal/pokecache"
)

type Client struct {
	Cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		Cache:      pokecache.NewCache(cacheInterval),
		httpClient: http.Client{Timeout: timeout},
	}
}
