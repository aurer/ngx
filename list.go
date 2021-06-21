package main

import (
	"fmt"
	"os"
	"path"
)

var List = Command {
	Name: "list",
	Aliases: []string{"ls", "l"},
	Description: "List all available configs",
	Run: func(args []string) error {
		configs := configFiles()
		configs.print()
		return nil
	},
}

func filesIn(dir string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return []os.DirEntry{}, err
	}

	var configFiles []os.DirEntry
	for _,file := range files {
		if !file.IsDir() {
			configFiles = append(configFiles, file)
		}
	}

	return configFiles, nil
}

func configFiles() Configs {
	available_files, err := filesIn(path_available)
		if err != nil {
			fmt.Printf("problem accessing %s\n", path_available)
			return Configs{}
		}

		enabled_files, err := filesIn(path_enabled)
		if err != nil {
			fmt.Printf("problem accessing %s\n", path_enabled)
			return Configs{}
		}

		var configs Configs
		for _,a := range available_files {
			enabled := false
			for _,e := range enabled_files {
				if e.Name() == a.Name() {
					enabled = true
				}
			}

			fileinfo,_ := a.Info()
			configs.add(Config{
				Name: a.Name(),
				Path: path.Join(path_available, a.Name()),
				Perm: fileinfo.Mode().Perm(),
				Enabled: enabled,
			})
		}

		return configs
}