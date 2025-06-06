// Package pokeapi provides functionality to interact with the PokeAPI
package pokeapi

// RespShallowLocations represents the response structure from the PokeAPI location-area endpoint
// It contains pagination information and a list of location areas
type RespShallowLocations struct {
	// Count is the total number of location areas available
	Count int `json:"count"`
	// Next is a pointer to the URL for the next page of results
	// It's a pointer because it can be null when there are no more pages
	Next *string `json:"next"`
	// Previous is a pointer to the URL for the previous page of results
	// It's a pointer because it can be null when on the first page
	Previous *string `json:"previous"`
	// Results is a slice of location areas, each containing a name and URL
	Results []struct {
		// Name is the name of the location area
		Name string `json:"name"`
		// URL is the API endpoint for this specific location area
		URL string `json:"url"`
	} `json:"results"`
}
