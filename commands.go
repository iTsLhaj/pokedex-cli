package main

import (
	"fmt"
	"io"
	"os"
)

const (
	helpMessage = `
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`

	baseUrl = "https://pokeapi.co/api/v2/"
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
}
