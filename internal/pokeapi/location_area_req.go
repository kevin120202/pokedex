package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kevin120202/pokedex/internal/pokecache"
)

func (c *Client) ListLocationAreas(locationUrl *string, cache *pokecache.Cache) (LocationAreasResp, error) {
	urlPath := "/location"
	fullUrl := baseUrl + urlPath

	if locationUrl != nil {
		// Check the cache for existing data
		if data, found := cache.Get(*locationUrl); found {
			var cachedLocation LocationAreasResp
			err := json.Unmarshal(data, &cachedLocation)
			if err != nil {
				return LocationAreasResp{}, err
			}
			return cachedLocation, nil
		}

		// If not in the cache, create a req
		fullUrl = *locationUrl
	}

	// Make the HTTP request if data is not found
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, res.Body)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	cache.Add(fullUrl, body)

	var location LocationAreasResp
	err = json.Unmarshal(body, &location)
	if err != nil {
		fmt.Println(err)
	}

	return location, nil
}
