package main

import (
	"io/ioutil"

	"github.com/antonlindstrom/valiant/config"
	"github.com/urfave/cli"
)

var cmdGenerate = cli.Command{
	Name:    "generate",
	Aliases: []string{"g", "gen"},
	Usage:   "generate example test file",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "test-file",
			Usage: "Full path to the destination, example: tests/00_example.yml (required)",
		},
	},
	Action: func(c *cli.Context) error {
		if c.String("test-file") == "" {
			return cli.NewExitError("--test-file is required", 1)
		}

		b, err := config.GenerateExample()
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		err = ioutil.WriteFile(c.String("test-file"), b, 0664)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		return nil
	},
}
