package main

import (
	cli "github.com/urfave/cli/v2"
)

func (suite *GLTestSuite) TestWhoAmI() {
	suite.IsType(&cli.Command{}, WhoAmI())
	suite.IsType((cli.ActionFunc)(nil), WhoAmI().Action)
	err := WhoAmI().Action(nil)
	suite.NoError(err)
}
