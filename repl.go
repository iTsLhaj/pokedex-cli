package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)



type (

	cliCommand struct {
		name string
		description string
		callback func() error
	}

	cliCmdRegister map[string]cliCommand

	ANSIColor string

)

const (
	ANSICLR_GREEN ANSIColor = "\033[32m"
	ANSICLR_BGREEN ANSIColor = "\033[1;32m"
	ANSICLR_RESET ANSIColor = "\033[0m"
)

func cleanInput(text string) []string {
	trimmedInput := strings.Trim(text, " \n\t\f\r\v")

	if trimmedInput == "" {
		return []string{}
	}

	words := strings.Split(trimmedInput, " ")
	var clean []string = make([]string, 0)
	
	for _, w := range words {
		clean = append(clean, strings.ToLower(strings.Trim(w, " \t")))
	}
	
	return clean
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	helpMessage := `
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`
	fmt.Print(helpMessage)
	return nil
}

func registerCommand(cmdReg cliCmdRegister, cmdName string, cmd cliCommand) error {
	_, ok := cmdReg[cmdName]
	if ok {
		return fmt.Errorf("command %s already exist", cmdName)
	}
	cmdReg[cmdName] = cmd
	return nil
}

func initCommands(commands cliCmdRegister) {
	var err error

	err = registerCommand(commands, "exit", cliCommand{
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = registerCommand(commands, "help", cliCommand{
		name: "help",
		description: "Display a Help message",
		callback: commandHelp,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func PokedexLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	var commands cliCmdRegister = make(cliCmdRegister, 0)

	initCommands(commands)
	for {
		fmt.Printf("\n%sPokedex%s > ", ANSICLR_BGREEN, ANSICLR_RESET)
		eof := scanner.Scan()
		if !eof {
			fmt.Printf("\nQuitting ...")
			break
		}
		cmd := cleanInput(scanner.Text())
		if len(cmd) == 0 {
			continue
		}
		for cmdName, cmdHandle := range commands {
			fmt.Println(cmdName)
			if cmd[0] == cmdName {
				fmt.Println("...")
				cmdHandle.callback()
			}
		}
	}
}
