package main

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

// WhoAmI shows you the currently logged in user
func WhoAmI() *cli.Command {
	return &cli.Command{
		Name:    "whoami",
		Aliases: []string{"me"},
		Usage:   "show the currently logged in user",
		Action: func(c *cli.Context) error {
			if baseURL == "" {
				baseURL = DefaultURL
			}
			client, err := NewClient(token, &baseURL)
			if err != nil {
				return cli.Exit(fmt.Sprintf("failed to create NewClient: %s", err), 1)
			}

			user, err := client.GetCurrentProfile()
			if err != nil {
				return cli.Exit(fmt.Sprintf("failed to GetCurrentProfile: %s", err), 1)
			}

			fmt.Printf("Name: %s\n", user.Name)
			return nil
		},
	}
}
