package main

import (
	"errors"
	"fmt"
	"os"
)

const path_available string = "/etc/nginx/sites-available"
const path_enabled string = "/etc/nginx/sites-enabled"

func main() {
	args := os.Args[1:]
	commands := []Command{
		List,
		Enable,
		Disable,
		Configtest,
	}

	handleHelp(commands, args)
	err := runCommands(commands, args)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func handleHelp(commands []Command, args []string) {
	if len(args) == 0 || args[0] == "--help" || args[0] == "help" {
		ShowHelp(commands)
		os.Exit(0)
	}
}

func runCommands(commands []Command, args []string) error {
	command := args[0]

	for _,c := range commands {
		if c.Name == command {
			err := c.Run(args[1:])
			if err != nil {
				fmt.Println(err)
			}
			return nil
		}

		for _,a := range c.Aliases {
			if a == command {
				err := c.Run(args[1:])
				if err != nil {
					fmt.Println(err)
				}
				return nil
			}
		}
	}

	message := fmt.Sprintf("the command '%s' was not found", command)
	return errors.New(message)
}