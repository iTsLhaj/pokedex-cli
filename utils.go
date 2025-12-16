package main

import (
	"fmt"
	"strings"
)

const (
	ANSICLR_GREEN  ANSIColor = "\033[32m"
	ANSICLR_BGREEN ANSIColor = "\033[1;32m"
	ANSICLR_RESET  ANSIColor = "\033[0m"
)

func putPrompt(p string) {
	fmt.Printf(
		"\n%s%s%s > ",
		ANSICLR_BGREEN, p, ANSICLR_RESET,
	)
}

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

func registerCommand(cmdReg cliCmdRegister, cmdName string, cmd cliCommand) error {
	_, ok := cmdReg[cmdName]
	if ok {
		return fmt.Errorf("command %s already exist", cmdName)
	}
	cmdReg[cmdName] = cmd
	return nil
}

func makeQuery(key, value string) string {
	return fmt.Sprintf("%s=%s", key, value)
}

func pokeDexInspect(pookie PokeDexPokemonData) string {
	data := fmt.Sprintf(`Name: %s
Height: %d
Weight: %d
Stats:
`, pookie.Name, pookie.Height, pookie.Weight)

	for _, stat := range pookie.Stats {
		data += fmt.Sprintf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	data += "Types:\n"

	for _, typ := range pookie.Types {
		data += fmt.Sprintf(" - %s\n", typ.Type.Name)
	}

	return data
}
