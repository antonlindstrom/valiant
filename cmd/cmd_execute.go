package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var cmdExecute = cli.Command{
	Name:    "execute",
	Aliases: []string{"e"},
	Usage:   "execute tests against address",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "address",
			Usage: "HTTP address to upstream server to test (required)",
		},
		cli.StringFlag{
			Name:  "test-directory",
			Value: "tests",
			Usage: "Path to directory containing tests",
		},
	},
	Action: func(c *cli.Context) {
		if c.String("address") == "" {
			fmt.Println("--address is required")
			os.Exit(1)
		}

		err := runTests(c.String("address"), c.String("test-directory"))
		if err == errTestFailure {
			os.Exit(1)
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
