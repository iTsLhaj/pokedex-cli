package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func NewPokeDexClient() *PokeDexClient {
	return &PokeDexClient{
		currentOffset: -pokeDexLimitOffset,
		nextOffset:    0,
	}
}

func (client *PokeDexClient) fetchLocations(httpClient *http.Client, offset, limit int) (data PokeDexLocationsData, err error) {
	endpoint := strings.Join(
		[]string{
			"location-area?",
			makeQuery("offset", strconv.Itoa(offset)),
			makeQuery("limit", strconv.Itoa(limit)),
		}, "&")
	url := strings.Join([]string{baseUrl, endpoint}, "/")

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
