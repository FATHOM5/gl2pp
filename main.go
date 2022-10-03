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
)

// Version returns the SemVer for this app.
func Version() string {
	return "v0.0.1"
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
			Value:       "https://gitlab.fathom5.work",
			Destination: &baseURL,
		},
	}

	// add commands in commends.go
	app.Commands = []*cli.Command{
		WhoAmI(),
	}

	app.Run(os.Args)
}
