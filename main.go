package main

import (
	"os"

	cli "github.com/urfave/cli/v2"
)

var (
	conf   Config
	logger *Logger

	app = cli.NewApp()
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
	app.Flags = []cli.Flag{}

	// add commands in commends.go
	app.Commands = []*cli.Command{
		VersionCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err, "failed to Run")
	}
}
