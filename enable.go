package main

import (
	"errors"
	"fmt"
	"os"
	"path"
)

var Enable = Command {
	Name: "enable",
	Aliases: []string{"en"},
	Description: "Enable a config",
	Run: func(args []string) error {
		if len(args) == 0 {
			return errors.New("please provide the name of a config")
		}

		configs := configFiles()
		foundConfig := false
		var selectedConfig Config

		for _,c := range configs.Configs {
			if c.Name == args[0] {
				foundConfig = true
				selectedConfig = c
			}
		}

		if !foundConfig {
			return fmt.Errorf("could not find a config matching '%s'", args[0])
		}

		if selectedConfig.Enabled {
			return fmt.Errorf("%s is already enabled", args[0])
		}

		source := path.Join(path_available, selectedConfig.Name)
		dest := path.Join(path_enabled, selectedConfig.Name)
		err := os.Symlink(source, dest)
		if err != nil {
			return err
		}

		fmt.Printf("enabled %s\n", selectedConfig.Name)
		return nil
	},
}