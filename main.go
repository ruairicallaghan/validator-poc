package main

import (
	"log"
	"validator-refactor/cmd"
)

func main() {
	root := cmd.NewRootCommand()
	root.AddCommand(
		cmd.NewCloudRunCommand(),
		cmd.NewDeploymentCommand(),
	)

	if err := root.Execute(); err != nil {
		log.Fatalf("%s\n", err)
	}
}
