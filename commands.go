package main

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

// VersionCommand prints out the version
func VersionCommand() *cli.Command {
	return &cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "returns version for this app",
		Action: func(c *cli.Context) error {
			logger.Info(fmt.Sprintf("Version: %s", Version()))
			return nil
		},
	}
}
