package main

import (
	"fmt"
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