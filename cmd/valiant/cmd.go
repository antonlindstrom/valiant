package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "valiant"
	app.Usage = "Validate request and responses to a web site."

	app.Commands = []cli.Command{
		cmdExecute,
		cmdUpdate,
		cmdGenerate,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute command: %s", err)
		os.Exit(1)
	}
}
