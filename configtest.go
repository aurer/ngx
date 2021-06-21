package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Configtest uses the 'nginx' utility to test the validity enabled config files
var Configtest = Command {
	Name: "configtest",
	Aliases: []string{"test"},
	Description: "Test the config is valid",
	Run: func(args []string) error {
		// Check the nginx utility exists
		path, err := exec.LookPath("nginx")
		if err != nil {
			return err
		}

		// Create the command
		cmd := exec.Command(path, "-t")

		// Setup variables to store output
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		// Run the command
		err = cmd.Run()
		if err != nil || stderr.String() != "" {
			return fmt.Errorf(stderr.String())
		}

		// Print the result
		fmt.Println(stdout.String())
		return nil
	},
}