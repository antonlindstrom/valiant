package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"
	"github.com/antonlindstrom/valiant/config"
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
	Action: func(c *cli.Context) {
		if c.String("test-file") == "" {
			fmt.Println("--test-file is required")
			os.Exit(1)
		}

		b, err := config.GenerateExample()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = ioutil.WriteFile(c.String("test-file"), b, 0664)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
