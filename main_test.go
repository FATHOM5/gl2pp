package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli"
)

type SkeletonCLITestSuite struct {
	suite.Suite
	Logger *Logger
	Conf   Config
	Ctx    *cli.Context
}

func (suite *SkeletonCLITestSuite) SetupTest() {
	suite.Conf, _ = NewConfig()
	suite.Logger = NewLogger(suite.Conf)
	suite.Ctx = &cli.Context{}
}

func (suite *SkeletonCLITestSuite) TestVersion() {
	suite.NotEqual("", Version())
}

func TestSkeletonCLITestSuite(t *testing.T) {
	suite.Run(t, new(SkeletonCLITestSuite))
}
