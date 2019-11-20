package main

import (
	"github.com/codegangsta/cli"
)

var cmdExecute = &cli.Command{
	Name:    "execute",
	Aliases: []string{"e"},
	Usage:   "execute tests against address",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "address",
			Usage: "HTTP address to upstream server to test (required)",
		},
		&cli.StringFlag{
			Name:  "test-directory",
			Value: "tests",
			Usage: "Path to directory containing tests",
		},
	},
	Action: func(c *cli.Context) error {
		if c.String("address") == "" {
			return cli.NewExitError("--address is required", 1)
		}

		err := runTests(c.String("address"), c.String("test-directory"))
		if err == errTestFailure {
			return cli.NewExitError("test exited with a failure", 1)
		}

		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		return nil
	},
}
