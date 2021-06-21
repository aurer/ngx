package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type Config struct {
	Name string
	Path string
	Perm os.FileMode
	Enabled bool
}

type Configs struct {
	Configs []Config
}

func (c *Configs) print() {
	maxNameLength := 0
	for _,config := range c.Configs {
		l := len(config.Name)
		if l > maxNameLength {
			maxNameLength = l
		}
	}

	for _,config := range c.Configs {
		l := len(config.Name)
		fmt.Print(config.Name)
		fmt.Print(strings.Repeat(" ", maxNameLength - l))
		fmt.Print(strings.Repeat(" ", 5))
		if config.Enabled {
			fmt.Print(color.GreenString("enabled"))
		} else {
			fmt.Print(color.RedString("disabled"))
		}
		fmt.Print("\n")
	}
}

func (c *Configs) add(config Config) {
	c.Configs = append(c.Configs, config)
}