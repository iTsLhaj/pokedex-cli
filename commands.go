package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

var client *PokeDexClient = NewPokeDexClient()

const (
	helpMessage = `Welcome to the Pokedex!
Usage:

map: Displays the next N location areas in the Pokemon world
mapb: (map back) Displays the previous N locations, similar to *map*
explore: list all pokemons located in *area*
catch: catch a pokemon
inspect: show details about an owned pokemon
help: Displays a help message
exit: Exit the Pokedex
`
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

func commandCatch(args []string) error {
	if len(args) == 0 {
		return errors.New("please provide a name, e.g: catch <pokemonName>")
	}
	if len(args) > 1 {
		return errors.New("one name is expected, got too many")
	}

	pookieData, err := client.fetchPookieData(&http.Client{}, args[0])
	if err != nil {
		return err
	}

	for _, pookie := range client.ownedPookies {
		if pookie.Name == args[0] {
			fmt.Printf("%s already caught!", pookie.Name)
			return nil
		}
	}

	chance := ((float64(pookieData.BaseExperience) - 400.0) / 400.0) * -1.0

	fmt.Printf("Throwing a Pokeball at %s...\n", pookieData.Name)
	if rand.Float64() < chance {
		/// caught!
		fmt.Printf("%s was caught!\n", pookieData.Name)
		client.ownedPookies = append(client.ownedPookies, pookieData)
	} else {
		/// escaped!
		fmt.Printf("%s escaped!\n", pookieData.Name)
	}
	return nil
}

func commandInspect(args []string) error {
	if len(args) == 0 {
		return errors.New("please provide a pokemonName")
	}
	if len(args) > 1 {
		return errors.New("one pokemonName is expected, got too many")
	}

	for _, pookie := range client.ownedPookies {
		if pookie.Name == args[0] {
			content := pokeDexInspect(pookie)
			fmt.Print(content)
			return nil
		}
	}

	fmt.Println("you have not caught that pokemon")
	return nil
}
