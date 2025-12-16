package main

import (
	"fmt"
	"io"
	"os"
)

var client *PokeDexClient = NewPokeDexClient()

const (
	helpMessage = `Welcome to the Pokedex!
Usage:

map: Displays the next N location areas in the Pokemon world
mapb: (map back) Displays the previous N locations, similar to *map*
help: Displays a help message
exit: Exit the Pokedex`
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelpMock(writer io.Writer) error {
	_, err := fmt.Fprint(writer, helpMessage)
	return err
}

func commandHelp() error {
	return commandHelpMock(os.Stdout)
}

func commandMap() error {
	locs, err := client.getNextLocations()
	if err != nil {
		return err
	}
	for _, loc := range locs {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapBack() error {
	locs, err := client.getPreviousLocations()
	if err != nil {
		return err
	}
	for _, loc := range locs {
		fmt.Println(loc.Name)
	}
	return nil
}
