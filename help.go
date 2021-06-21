package main

import (
	"fmt"
)

func ShowHelp(commands []Command) {
	fmt.Println("NGX is a command line tool for managing Nginx")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println()
	for _, c := range commands {
		c.Describe()
	}
	fmt.Println()
}