package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func initCommands(commands cliCmdRegister) {
	var err error

	err = registerCommand(commands, "exit", cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = registerCommand(commands, "help", cliCommand{
		name:        "help",
		description: "Display a Help message",
		callback:    commandHelp,
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
			fmt.Println(cmdName)
			if cmd[0] == cmdName {
				fmt.Println("...")
				err := cmdHandle.callback()
				if err != nil {
					fmt.Printf(" <%s> Failed: %s\n", cmdName, err)
				}
			}
		}
	}
}
