package main

import (
	"fmt"
	"os"
)

type Command struct {
	Name string
	Aliases []string
	Description string
	Run func(args []string) error
}

func (c Command) Describe() {
	name := c.Name;
	for _,a := range c.Aliases {
		name += ", " + a
	}
	out := fmt.Sprintf("\t%s\t\t%s", name, c.Description)
	fmt.Println(out)
}

var commands []Command

func main() {
	args := os.Args[1:]

	helpCommand := Command{
		Name: "test",
		Aliases: []string{"t"},
		Description: "Test if things work",
		Run: func(args []string) error {
			fmt.Println(args)
			return nil
		},
	}

	commands = append(commands, helpCommand)

	if len(args) == 0 || args[0] == "--help" || args[0] == "help" {
		showHelp()
		os.Exit(0)
	}

	command := args[0]

	for _,c := range commands {
		if c.Name == command {
			c.Run(args[1:])
			break
		}

		for _,a := range c.Aliases {
			if a == command {
				c.Run(args[1:])
				break
			}
		}
	}
}

func showHelp() {
	fmt.Println("NGX is a command line tool for managing Nginx")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println()
	for _, c := range commands {
		c.Describe()
	}
	fmt.Println()
}