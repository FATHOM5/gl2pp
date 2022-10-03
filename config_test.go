package main

func (suite *GLTestSuite) TestNewConfig() {
	conf, err := NewConfig()
	suite.NoError(err)
	suite.IsType(Config{}, conf)
}

func (suite *GLTestSuite) TestConfig() {
	suite.NotEmpty(suite.Conf.AppName)
	suite.Equal("test", suite.Conf.AppEnv)
}
