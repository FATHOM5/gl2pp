package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli/v2"
)

type GLTestSuite struct {
	suite.Suite
	Logger      *Logger
	Conf        Config
	Ctx         *cli.Context
	GitlabToken string
	BaseURL     string
	Client      *Client
}

func (suite *GLTestSuite) SetupTest() {
	suite.Conf, _ = NewConfig()
	suite.Logger = NewLogger(suite.Conf)
	suite.Ctx = &cli.Context{}
	suite.GitlabToken = os.Getenv("GITLAB_CI_TOKEN")
	suite.BaseURL = "https://gitlab.fathom5.work"

	client, err := NewClient(suite.GitlabToken, nil)
	if err != nil {
		log.Fatalf("failed to create NewClient: %s", err)
	}
	suite.Client = client

	// set default values for global flags
	baseURL = suite.BaseURL
	token = suite.GitlabToken
}

func (suite *GLTestSuite) TestVersion() {
	suite.NotEqual("", Version())
}

func TestGLTestSuite(t *testing.T) {
	suite.Run(t, new(GLTestSuite))
}
