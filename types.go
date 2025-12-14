package main

type (
	cliCommand struct {
		name        string
		description string
		callback    func() error
	}

	cliCmdRegister map[string]cliCommand

	ANSIColor string

	pokeDexLocation struct {
		name string
		url  string
	}

	pokeDexClient struct {
		locations []pokeDexLocation
	}
)
