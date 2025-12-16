package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/kenzo/pokedexcli/internal/pokecache"
)

func NewPokeDexClient() *PokeDexClient {
	return &PokeDexClient{
		baseURL:       "https://pokeapi.co/api/v2/",
		cache:         pokecache.NewCache(),
		currentOffset: -pokeDexLimitOffset,
		nextOffset:    0,
	}
}

func (client *PokeDexClient) fetchPokiesAt(httpClient *http.Client, area string) (pookies []PokeDexPokemon, err error) {
	var data PokeDexPokemonEncounters

	pookies, err = client.isPDPEsEntryCached(area)
	if err == nil {
		return
	}

	endpoint := strings.Join(
		[]string{
			client.baseURL,
			"location-area",
			area,
		}, "/")
	var req *http.Request
	req, err = http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return
	}

	var res *http.Response
	res, err = httpClient.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode == 404 {
		err = errors.New("area not found")
		return
	}
	if res.StatusCode != http.StatusOK {
		err = errors.New("request to the pokedex API failed")
		return
	}

	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	client.cache.Cache.Add(area, body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}

	pookies = make([]PokeDexPokemon, len(data.PokemonEncounters))
	for _, encounteredPookie := range data.PokemonEncounters {
		pookies = append(pookies, encounteredPookie.Pokemon)
	}

	return
}

func (client *PokeDexClient) fetchLocations(httpClient *http.Client, offset, limit int) (data PokeDexLocationsData, err error) {
	data, err = client.isPDLDEntryCached(offset)
	if err == nil {
		return
	}

	endpoint := strings.Join(
		[]string{
			"location-area?",
			makeQuery("offset", strconv.Itoa(offset)),
			makeQuery("limit", strconv.Itoa(limit)),
		}, "&")
	url := strings.Join([]string{client.baseURL, endpoint}, "/")

	var req *http.Request
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	var res *http.Response
	res, err = httpClient.Do(req)
	if err != nil {
		return
	}

	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	client.cache.Cache.Add(strconv.Itoa(offset), body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}

	return
}

func (client *PokeDexClient) checkOffsets() error {
	if client.currentOffset < 0 || client.nextOffset < 0 {
		return errors.New("offset cannot be negative")
	}
	if client.currentOffset >= client.nextOffset {
		return errors.New("offset cannot be greater than or equal to next")
	}
	return nil
}

func (client *PokeDexClient) getPreviousLocations() ([]PokeDexLocation, error) {
	if client.currentOffset == 0 {
		return []PokeDexLocation{}, errors.New("current offset is 0, can't fetch previous locations")
	}
	client.currentOffset -= pokeDexLimitOffset
	client.nextOffset -= pokeDexLimitOffset
	locData, err := client.fetchLocations(&http.Client{}, client.currentOffset, pokeDexLimitOffset)
	if err != nil {
		return []PokeDexLocation{}, err
	}
	return locData.Locations, nil
}

func (client *PokeDexClient) getNextLocations() ([]PokeDexLocation, error) {
	client.currentOffset += pokeDexLimitOffset
	client.nextOffset += pokeDexLimitOffset
	locData, err := client.fetchLocations(&http.Client{}, client.currentOffset, pokeDexLimitOffset)
	if err != nil {
		return []PokeDexLocation{}, err
	}
	return locData.Locations, nil
}

func (client *PokeDexClient) isPDLDEntryCached(offset int) (PokeDexLocationsData, error) {
	var data PokeDexLocationsData

	body, ok := client.cache.Cache.Get(strconv.Itoa(offset))
	if !ok {
		return PokeDexLocationsData{}, errors.New("cache entry not found")
	}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return PokeDexLocationsData{}, err
	}
	return data, nil
}

func (client *PokeDexClient) isPDPEsEntryCached(area string) ([]PokeDexPokemon, error) {
	var data PokeDexPokemonEncounters

	body, ok := client.cache.Cache.Get(area)
	if !ok {
		return []PokeDexPokemon{}, errors.New("cache entry not found")
	}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return []PokeDexPokemon{}, err
	}

	pokies := make([]PokeDexPokemon, len(data.PokemonEncounters))
	for _, encounteredPookie := range data.PokemonEncounters {
		pokies = append(pokies, encounteredPookie.Pokemon)
	}

	return pokies, nil
}
