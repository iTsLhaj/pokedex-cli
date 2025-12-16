package main

import "github.com/kenzo/pokedexcli/internal/pokecache"

const pokeDexLimitOffset = 20

type (
	cliCommand struct {
		name        string
		description string
		callback    func() error
	}

	cliCmdRegister map[string]cliCommand

	ANSIColor string

	PokeDexLocation struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	PokeDexLocationsData struct {
		LocationsCount int               `json:"count"`
		Next           string            `json:"next"`
		Previous       any               `json:"previous"`
		Locations      []PokeDexLocation `json:"results"`
	}

	PokeDexClient struct {
		baseURL       string
		cache         *pokecache.ConfigurableCache
		currentOffset int
		nextOffset    int
	}
)
