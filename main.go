package main

import (
	"os"

	cli "github.com/urfave/cli/v2"
)

var (
	conf   Config
	logger *Logger

	app = cli.NewApp()

	// global flags
	baseURL string
	token   string

	//other flags
	gid      string
	iid      string
	filename string
)

// SemVer is a string reprsentation of the Semantic Version. This will be injected
// during build time using the value from `git describe`.
var SemVer = "v0.0.0-development"

// Version returns the SemVer for this app.
func Version() string {
	return SemVer
}

func init() {
	conf, err := NewConfig()
	if err != nil {
		panic(err)
	}

	logger = NewLogger(conf)

	app.Name = conf.AppName
	app.Usage = conf.AppDesc
	app.Version = Version()
}

func main() {
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "base-url",
			Aliases:     []string{"url"},
			Usage:       "set the base url of the gitlab instance you want to use",
			EnvVars:     []string{"GITLAB_BASE_URL"},
			Value:       "https://github.com",
			Destination: &baseURL,
		},
		&cli.StringFlag{
			Name:        "token",
			Aliases:     []string{"t"},
			Usage:       "your personal access token; used to authenticate against the gitlab api",
			EnvVars:     []string{"GITLAB_TOKEN"},
			Destination: &token,
		},
	}

	app.Commands = []*cli.Command{
		WhoAmI(),
		ListGroups(),
		ListGroupIterations(),
		ListGroupIssues(),
	}

	app.Run(os.Args)
}
