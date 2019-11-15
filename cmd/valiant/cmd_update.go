package main

import (
	"fmt"
	"io/ioutil"

	"github.com/codegangsta/cli"
)

var cmdUpdate = &cli.Command{
	Name:    "update",
	Aliases: []string{"u"},
	Usage:   "update test with upstream response",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "address",
			Usage: "HTTP address to upstream server to test (required)",
		},
		&cli.StringFlag{
			Name:  "test-file",
			Usage: "Full path to the test to update, example: tests/00_example.yml (required)",
		},
	},
	Action: func(c *cli.Context) error {
		if c.String("address") == "" {
			return cli.NewExitError("--address is required", 1)
		}

		if c.String("test-file") == "" {
			return cli.NewExitError("--test-file is required", 1)
		}

		spec, err := parseFile(c.String("test-file"))
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		resp, err := spec.SendRequest(c.String("address"))
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		spec.Response = *resp

		b, err := spec.Update()
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		err = ioutil.WriteFile(c.String("test-file"), b, 0664)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		fmt.Printf("%s updated successfully\n", c.String("test-file"))

		return nil
	},
}
