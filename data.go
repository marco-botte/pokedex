package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"pokedex/internal/pokecache"
)

type Stat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type Type struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

type Pokemon struct {
	Name   string `json:"-"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []Stat `json:"stats"`
	Types  []Type `json:"types"`
}

type PokemonExperience struct {
	Experience int `json:"base_experience"`
}

type LocationResponse struct {
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
type LocationRef struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocAreaResponse struct {
	Next     *string       `json:"next"`
	Previous *string       `json:"previous"`
	Results  []LocationRef `json:"results"`
}

func getData[T any](url string, cache *pokecache.Cache) (*T, error) {
	body, cached := cache.Get(url)
	if !cached {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
			return nil, err
		}
		cache.Add(url, body)
	}
	var result T
	err := json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &result, nil
}
