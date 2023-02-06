package main

import (
	"fmt"
	"strconv"
	"time"

	cli "github.com/urfave/cli/v2"
)

// WhoAmI shows you the currently logged in user
func WhoAmI() *cli.Command {
	return &cli.Command{
		Name:      "whoami",
		Aliases:   []string{"me"},
		ArgsUsage: "gl2pp whoami",
		Usage:     "show the currently logged in user",
		Action: func(c *cli.Context) error {
			start := time.Now()
			if c != nil {
				fmt.Printf("%s\n\n", c.Command.Name)
			}

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
			fmt.Printf("Email: %s\n", user.Email)
			fmt.Printf("Username: %s\n", user.Username)
			fmt.Printf("Profile: %s\n", user.Profile)
			fmt.Printf("SSHKey: %s\n", user.SSHKey)
			fmt.Printf("GPGKey: \n%s\n", user.GPGKey)
			end := time.Now()

			fmt.Printf("%s", end.Sub(start))
			return nil
		},
	}
}

// ListGroups lists the groups the current user has access to on Gitlab.
func ListGroups() *cli.Command {
	return &cli.Command{
		Name:      "list-groups",
		Aliases:   []string{"grp"},
		ArgsUsage: "gl2pp list-groups",
		Usage:     "Show groups for current logged in user",
		Action: func(c *cli.Context) error {
			if baseURL == "" {
				baseURL = DefaultURL
			}
			client, err := NewClient(token, &baseURL)
			if err != nil {
				return cli.Exit(fmt.Sprintf("failed to create NewClient: %s", err), 1)
			}

			groups, err := client.ListGroups()
			if err != nil {
				return cli.Exit(fmt.Sprintf("failed to list current groups: %s", err), 1)

			}
			PrintGroupsTable(groups)

			return nil
		},
	}
}

// ListGroupIterations lists sprints for the given group ID.
func ListGroupIterations() *cli.Command {
	return &cli.Command{
		Name:      "list-group-iterations",
		Aliases:   []string{"it"},
		ArgsUsage: "gl2pp it --group-id=<value>",
		Usage:     "Show group iterations for the corresponding group-id",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "group-id",
				Aliases:     []string{"gid"},
				Usage:       "This flag supplies the group id from GitLab to the get iterations",
				Destination: &gid,
			},
		},
		Action: func(c *cli.Context) error {
			if baseURL == "" {
				baseURL = DefaultURL
			}
			if gid == "" {
				return cli.Exit("Failed from group-id missing, please provide group-id to cli", 1)
			}
			client, err := NewClient(token, &baseURL)
			if err != nil {
				return cli.Exit(fmt.Sprintf("failed to create NewClient: %s", err), 1)
			}

			iterations, err := client.ListGroupIterations(gid)
			if err != nil {
				return cli.Exit(fmt.Sprintf("failed to list group iterations: %s", err), 1)
			}

			PrintItterationsTable(iterations)

			return nil
		},
	}
}

// ListGroupIssues lists the issues for the given group ID and iteration ID.
func ListGroupIssues() *cli.Command {
	return &cli.Command{
		Name:      "list-group-issues",
		Aliases:   []string{"iss"},
		ArgsUsage: "gl2pp iss --group-id=<value>  --iteration-id=<value>  --output=<value>",
		Usage:     "Show group issues for the current selected group-id/iteration and give a filename for the output file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "group-id",
				Aliases:     []string{"gid"},
				Usage:       "This flag supplies the group-id to the list group issues as an arguement",
				Destination: &gid,
			},
			&cli.StringFlag{
				Name:        "iteration-id",
				Aliases:     []string{"iid"},
				Usage:       "This flag supplies the iteration-id to the list group issues as an arguement",
				Destination: &iid,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "Provides filename for output",
				Destination: &filename,
			},
		},
		Action: func(c *cli.Context) error {
			if baseURL == "" {
				baseURL = DefaultURL
			}

			client, err := NewClient(token, &baseURL)
			if err != nil {
				return cli.Exit(fmt.Sprintf("failed to create NewClient: %s", err), 1)
			}

			if gid == "" {
				return cli.Exit("Failed from token missing, please provide gid to cli", 1)
			}

			var iidplaceholder int
			iidplaceholder, err = strconv.Atoi((iid))
			if err != nil {
				return cli.Exit(fmt.Sprintf("failed to convert iid.  Iid must be type integer: %s", err), 1)
			}

			issues, err := client.ListGroupIssues(gid, iidplaceholder)
			if err != nil {
				return cli.Exit(fmt.Sprintf("Failed to list group issues: %s", err), 1)
			}

			if filename != "" {
				outissues := translate(issues)
				err = Exporttocsv(outissues, filename)
				if err != nil {
					return cli.Exit(fmt.Sprintf("Failed to save file: %s", err), 1)
				}
				fmt.Printf("File with issues saved succesfully to %s", filename)
				return nil
			}
			PrintIssuesTable(issues)
			return nil
		},
	}
}
