package main

import (
	"net/http"
	"strings"
)

func (client *pokeDexClient) fetchLocations(httpClient *http.Client) error {
	endpoint := "location"
	req, err := http.NewRequest("GET", strings.Join([]string{baseUrl, endpoint}, "/"), nil)
}
