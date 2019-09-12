package main

import (
	"fmt"
	"golb/commands"
	"golb/settings"
	"os"
)

const BUILD_COMMAND = "build"
const DEV_WEBSERVER_COMMAND = "develop"

var availableCommands []string = getAvailableCommands()

func main() {
	command := os.Args[1]

	if isValidCommand(command) {
		defer fmt.Printf("\n\nEverything done! Goodbye.\n\n")
		settings := settings.Instance()

		switch command {
		case BUILD_COMMAND:
			commands.Build()
			break
		case DEV_WEBSERVER_COMMAND:
			commands.Develop(settings)
			break
		}
	} else {
		fmt.Printf("%s: command not found.\n", command)
	}

}

func getAvailableCommands() []string {
	return []string{BUILD_COMMAND, DEV_WEBSERVER_COMMAND}
}

func isValidCommand(command string) bool {
	for _, n := range availableCommands {
		if command == n {
			return true
		}
	}
	return false
}
