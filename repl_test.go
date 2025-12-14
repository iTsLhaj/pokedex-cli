package main

import (
	"reflect"
	"testing"
)



func TestCleanInput(t *testing.T) {
	cases := []struct{
		input string
		expected []string
	}{
		{ input: "Charmander Bulbasaur PIKACHU", expected: []string{ "charmander", "bulbasaur", "pikachu" } },
	}

	for _, tc := range cases {
		got := cleanInput(tc.input)
		if !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("got %v expected %v", got, tc.expected)
		}
	}
}
