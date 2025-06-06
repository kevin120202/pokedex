package main

import (
	"errors"
	"fmt"
)

// commandMapf handles the "map" command to display the next page of locations
// It fetches location data from the PokeAPI and updates the pagination URLs
func commandMapf(cfg *config, args ...string) error {
	// Fetch the next page of locations using the current nextLocationsURL
	// This makes an HTTP request to the PokeAPI
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	// Update the config with the new pagination URLs
	// These will be used for the next map/mapb commands
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	// Print each location name from the response
	// This displays the list of locations to the user
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// commandMapb handles the "mapb" command to display the previous page of locations
// It's similar to commandMapf but uses the previous URL instead
func commandMapb(cfg *config, args ...string) error {
	// Check if we're on the first page (no previous URL)
	// Return an error if there's no previous page to go back to
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	// Fetch the previous page of locations using the current prevLocationsURL
	// This makes an HTTP request to the PokeAPI
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	// Update the config with the new pagination URLs
	// These will be used for the next map/mapb commands
	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	// Print each location name from the response
	// This displays the list of locations to the user
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
