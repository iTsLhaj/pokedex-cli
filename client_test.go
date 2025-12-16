package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchLocations_UsesCache(t *testing.T) {
	mockJSON := `{
		"count": 1,
		"results": [
			{ "name": "canalave-city-area", "url": "https://pokeapi.co/api/v2/location-area/1/" }
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockJSON))
	}))
	defer server.Close()

	client := NewPokeDexClient()
	client.baseURL = server.URL
	httpClient := &http.Client{}

	data1, err := client.fetchLocations(httpClient, 0, 20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(data1.Locations) != 1 {
		t.Fatalf("expected 1 location, got %d", len(data1.Locations))
	}

	data2, err := client.fetchLocations(httpClient, 0, 20)
	if err != nil {
		t.Fatalf("unexpected error on cache hit: %v", err)
	}

	if data2.Locations[0].Name != "canalave-city-area" {
		t.Fatalf("unexpected location name: %s", data2.Locations[0].Name)
	}
}
