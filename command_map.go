package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func mapCallback() {
	res, err := http.Get("https://pokeapi.co/api/v2/location?limit=20&offset=20")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	var location Location
	err = json.Unmarshal(body, &location)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(location.Results); i++ {
		fmt.Printf("%d: %s\n", i+1, location.Results[i].Name)
	}
}
