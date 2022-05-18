package main

import (
	cli "github.com/urfave/cli/v2"
)

func (suite *SkeletonCLITestSuite) TestVersionCommand() {
	suite.IsType(&cli.Command{}, VersionCommand())
	suite.IsType((cli.ActionFunc)(nil), VersionCommand().Action)
}
