package main

import (
	"errors"
	"io"
)

func (suite *GLTestSuite) TestlogWriter() {
	suite.Implements((*io.Writer)(nil), logWriter("test"))
}

func (suite *GLTestSuite) TestNewLogger() {
	suite.IsType(&Logger{}, NewLogger(suite.Conf))
}

func (suite *GLTestSuite) TestAnd() {
	suite.IsType(&Logger{}, suite.Logger.And("", ""))
}

func (suite *GLTestSuite) TestInfo() {
	suite.IsType(&Logger{}, suite.Logger.Info(""))
}

func (suite *GLTestSuite) TestDebug() {
	suite.IsType(&Logger{}, suite.Logger.Debug(""))
}

func (suite *GLTestSuite) TestLoggerError() {
	suite.IsType(&Logger{}, suite.Logger.Error(errors.New(""), ""))
}
