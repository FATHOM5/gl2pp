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
		Name:    "whoami",
		Aliases: []string{"me"},
		Usage:   "show the currently logged in user",
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

func ListGroups() *cli.Command {
	return &cli.Command{
		Name:    "list-groups",
		Aliases: []string{"grp"},
		Usage:   "Show groups for current logged in user",
		Action: func(c *cli.Context) error {
			//start := time.Now()
			// if c != nil {
			// 	fmt.Printf("%s,\n, \n", c.Command.Name)
			// }

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

			for _, group := range groups {
				fmt.Printf("GroupName: %s\n", group.Name)
				fmt.Printf("GroupID: %d\n", group.ID)
				fmt.Printf("GroupWebURL: %s\n", group.WebURL)
			}

			return nil
		},
	}
}

func ListGroupIterations() *cli.Command {
	return &cli.Command{
		Name:    "list-group-iterations",
		Aliases: []string{"it"},
		Usage:   "Show group iterations for current logged in user",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "group-id",
				Aliases:     []string{"gid"},
				Usage:       "This flag supplies the group id from GitLab to the get iterations",
				Destination: &gid,
			},
		},
		Action: func(c *cli.Context) error {
			//start := time.Now()

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

			for _, iteration := range iterations {
				fmt.Printf("GroupIterationTitle: %s\n", iteration.Title)
				fmt.Printf("GroupIterationID: %d\n", iteration.ID)
				fmt.Printf("GroupIterationWebURL: %s\n", iteration.WebURL)
			}

			return nil
		},
	}
}

func ListGroupIssues() *cli.Command {
	return &cli.Command{
		Name:    "list-group-issues",
		Aliases: []string{"iss"},
		Usage:   "Show group issues for the current logged in user",
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
			//start := time.Now()

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
			for _, issue := range issues {
				fmt.Printf("GroupIssuesTitle: %s\n", issue.Title)
				// fmt.Printf("GroupIssuesID: %d\n", issue.ID)
				fmt.Printf("GroupIssuesIID: %d\n", issue.IID)
				// fmt.Printf("GroupIssuesDescription: %s\n", issue.Description)
				fmt.Printf("GroupIssuesProjectID: %d\n", issue.ProjectID)
				fmt.Printf("GroupIssuesEpic: %p\n", issue.Epic)
				fmt.Printf("GroupIssuesWeight: %d\n", issue.Weight)
				fmt.Printf("GroupIssuesWeburl: %s\n", issue.WebURL)
			}

			return nil
		},
	}
}
