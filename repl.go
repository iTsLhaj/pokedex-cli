package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var commandsList []struct {
	cmdName string
	cmd     cliCommand
} = []struct {
	cmdName string
	cmd     cliCommand
}{
	{
		cmdName: "exit",
		cmd: cliCommand{
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	},
	{
		cmdName: "help",
		cmd: cliCommand{
			name:        "help",
			description: "Display a Help message",
			callback:    commandHelp,
		},
	},
	{
		cmdName: "map",
		cmd: cliCommand{
			name:        "map",
			description: "Display World Map",
			callback:    commandMap,
		},
	},
	{
		cmdName: "mapb",
		cmd: cliCommand{
			name:        "mapb",
			description: "Display Previous World Map",
			callback:    commandMapBack,
		},
	},
	{
		cmdName: "explore",
		cmd: cliCommand{
			name:        "explore",
			description: "Display List of Pokemons in Area",
			callback:    commandExplore,
		},
	},
}

func initCommands(commands cliCmdRegister) {
	var err error

	for _, cmd := range commandsList {
		err = registerCommand(commands, cmd.cmdName, cmd.cmd)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func PokedexLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	var commands cliCmdRegister = make(cliCmdRegister, 0)

	initCommands(commands)
	for {
		commandFound := false

		putPrompt("Pokedex")
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
			if cmd[0] == cmdName {
				commandFound = true
				err := cmdHandle.callback(cmd[1:])
				if err != nil {
					fmt.Printf(" <%s> Failed: %s\n", cmdName, err)
				}
			}
		}

		if !commandFound {
			fmt.Printf("command `%s` not found\n", cmd[0])
		}
	}
}
