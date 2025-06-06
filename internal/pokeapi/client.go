package pokeapi

import (
	"net/http"
	"time"

	"github.com/kevin120202/pokedex/internal/pokecache"
)

// Client represents a PokeAPI client with a configured HTTP client
// It's used to make requests to the PokeAPI with proper timeout settings
type Client struct {
	cache pokecache.Cache
	// httpClient is the underlying HTTP client used to make requests
	// It's configured with a timeout to prevent hanging requests
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
