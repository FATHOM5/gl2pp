package main

func (suite *GLTestSuite) TestNewCLient() {
	client, err := NewClient(suite.GitlabToken, nil)
	suite.NoError(err)
	suite.NotNil(client)
	suite.Equal("https://gitlab.fathom5.work", client.URL)

}

func (suite *GLTestSuite) TestNewCLientCustomBaseURL() {
	url := "https://gitlab.erm.fathom5.work"
	client, err := NewClient(suite.GitlabToken, &url)
	suite.NoError(err)
	suite.NotNil(client)
	suite.Equal("https://gitlab.erm.fathom5.work", client.URL)
	DefaultURL = "https://gitlab.fathom5.work" // reset global variable
}

func (suite *GLTestSuite) TestGetCurrentProfile() {
	user, err := suite.Client.GetCurrentProfile()
	suite.NoError(err)
	suite.NotNil(user)
	suite.NotEmpty(user.Name)
}
