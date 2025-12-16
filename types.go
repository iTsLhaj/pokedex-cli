package main

import "github.com/kenzo/pokedexcli/internal/pokecache"

const pokeDexLimitOffset = 20

type (
	cliCommand struct {
		name        string
		description string
		callback    func([]string) error
	}

	cliCommandEntry struct {
		cmdName string
		cmd     cliCommand
	}

	cliCmdRegister map[string]cliCommand

	ANSIColor string

	PokeDexEntity struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	PokeDexLocation PokeDexEntity

	PokeDexLocationsData struct {
		LocationsCount int               `json:"count"`
		Next           string            `json:"next"`
		Previous       any               `json:"previous"`
		Locations      []PokeDexLocation `json:"results"`
	}

	PokeDexPokemon PokeDexEntity

	PokeDexPokemonEncounters struct {
		ID                int           `json:"id"`
		Location          PokeDexEntity `json:"location"`
		Name              string        `json:"name"`
		PokemonEncounters []struct {
			Pokemon        PokeDexPokemon `json:"pokemon"`
			VersionDetails []struct {
				EncounterDetails []struct {
					Chance          int           `json:"chance"`
					ConditionValues []any         `json:"condition_values"`
					MaxLevel        int           `json:"max_level"`
					Method          PokeDexEntity `json:"method"`
					MinLevel        int           `json:"min_level"`
				} `json:"encounter_details"`
				MaxChance int           `json:"max_chance"`
				Version   PokeDexEntity `json:"version"`
			} `json:"version_details"`
		} `json:"pokemon_encounters"`
	}

	PokeDexPDStat struct {
		BaseStat int           `json:"base_stat"`
		Effort   int           `json:"effort"`
		Stat     PokeDexEntity `json:"stat"`
	}

	PokeDexPDType struct {
		Slot int           `json:"slot"`
		Type PokeDexEntity `json:"type"`
	}

	PokeDexPokemonData struct {
		BaseExperience int             `json:"base_experience"`
		Weight         int             `json:"weight"`
		Height         int             `json:"height"`
		Name           string          `json:"name"`
		Stats          []PokeDexPDStat `json:"stats"`
		Types          []PokeDexPDType `json:"types"`
	}

	PokeDexClient struct {
		baseURL       string
		cache         *pokecache.ConfigurableCache
		currentOffset int
		nextOffset    int
		ownedPookies  []PokeDexPokemonData
	}
)
