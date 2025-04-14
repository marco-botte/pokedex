package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocAreaResponse struct {
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

func getLocs(url string) (*LocAreaResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return nil, err
	}
	var locs LocAreaResponse
	err = json.Unmarshal(body, &locs)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}
	return &locs, nil
}
