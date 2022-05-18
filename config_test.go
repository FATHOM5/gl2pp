package main

func (suite *SkeletonCLITestSuite) TestNewConfig() {
	conf, err := NewConfig()
	suite.NoError(err)
	suite.IsType(Config{}, conf)
}

func (suite *SkeletonCLITestSuite) TestConfig() {
	suite.NotEmpty(suite.Conf.AppName)
	suite.Equal("test", suite.Conf.AppEnv)
}
