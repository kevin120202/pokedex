package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + fmt.Sprintf("/location-area/%s", locationName)

	if val, ok := c.cache.Get(url); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}

		return locationResp, nil
	}

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	// Execute the request using the client's HTTP client
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	// Ensure the response body is closed after we're done with it
	defer resp.Body.Close()

	// Read the entire response body
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)
	return locationResp, nil
}
