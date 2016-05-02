package main

import (
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

	app.Run(os.Args)
}
