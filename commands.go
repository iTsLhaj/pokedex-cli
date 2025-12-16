package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var client *PokeDexClient = NewPokeDexClient()

const (
	helpMessage = `Welcome to the Pokedex!
Usage:

explore: list all pokemons located in *area*
map: Displays the next N location areas in the Pokemon world
mapb: (map back) Displays the previous N locations, similar to *map*
help: Displays a help message
exit: Exit the Pokedex`
)

func commandExit(args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelpMock(writer io.Writer) error {
	_, err := fmt.Fprint(writer, helpMessage)
	return err
}

func commandHelp(args []string) error {
	return commandHelpMock(os.Stdout)
}

func commandMap(args []string) error {
	locs, err := client.getNextLocations()
	if err != nil {
		return err
	}
	for _, loc := range locs {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapBack(args []string) error {
	locs, err := client.getPreviousLocations()
	if err != nil {
		return err
	}
	for _, loc := range locs {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(args []string) error {
	if len(args) == 0 {
		return errors.New("please provide an area, e.g: explore <area>")
	}
	if len(args) > 1 {
		return errors.New("one area is expected, got too many")
	}
	pookies, err := client.fetchPokiesAt(&http.Client{}, args[0])
	if err != nil {
		return err
	}
	for _, pookie := range pookies {
		if len(strings.Trim(pookie.Name, "\n\t ")) == 0 {
			continue
		}
		fmt.Println(pookie.Name)
	}
	return nil
}
