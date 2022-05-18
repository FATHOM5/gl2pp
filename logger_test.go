package main

import (
	"errors"
	"io"
)

func (suite *SkeletonCLITestSuite) TestlogWriter() {
	suite.Implements((*io.Writer)(nil), logWriter("test"))
}

func (suite *SkeletonCLITestSuite) TestNewLogger() {
	suite.IsType(&Logger{}, NewLogger(suite.Conf))
}

func (suite *SkeletonCLITestSuite) TestLoggerLogError() {
	err := errors.New("test error")
	suite.Error(suite.Logger.Error(err, ""))
}
