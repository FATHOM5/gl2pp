package main

import (
	gitlab "github.com/xanzy/go-gitlab"
)

// Client will manage all api interactions with Gitlab
type Client struct {
	URL   string
	Token string
	GL    *gitlab.Client
}

// DefaultURL of the gitlab instance we want to use.
var DefaultURL = "https://gitlab.fathom5.work"

// NewClient will return a initialized Client
func NewClient(token string, url *string) (*Client, error) {
	if url != nil {
		DefaultURL = *url
	}

	gl, err := gitlab.NewClient(token, gitlab.WithBaseURL(DefaultURL))
	if err != nil {
		return nil, err
	}

	client := &Client{
		URL:   DefaultURL,
		Token: token,
		GL:    gl,
	}

	return client, nil
}

// GetCurrentProfile returns the currently logged in user's profile.
func (c *Client) GetCurrentProfile() (*GitlabUser, error) {
	var (
		user    *gitlab.User
		gpgKeys []*gitlab.GPGKey
		gpgKey  = "No GPG Key"
		sshKeys []*gitlab.SSHKey
		sshKey  = "No SSH Key"
		err     error
	)

	user, _, err = c.GL.Users.CurrentUser()
	if err != nil {
		return nil, err
	}

	gpgKeys, _, err = c.GL.Users.ListGPGKeys()
	if err != nil {
		return nil, err
	}

	if len(gpgKeys) > 0 {
		gpgKey = gpgKeys[0].Key
	}

	sshKeys, _, err = c.GL.Users.ListSSHKeys()
	if err != nil {
		return nil, err
	}

	if len(sshKeys) > 0 {
		sshKey = sshKeys[0].Key
	}

	glUser := &GitlabUser{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Profile:  user.WebURL,
		GPGKey:   gpgKey,
		SSHKey:   sshKey,
	}

	return glUser, nil
}

func (c *Client) ListGroups() ([]*gitlab.Group, error) {
	var (
		grpopt gitlab.ListGroupsOptions
		gs     []*gitlab.Group
		err    error
	)
	gs, _, err = c.GL.Groups.ListGroups(&grpopt)
	return gs, err

}

var state *string

func (c *Client) ListGroupIterations(gid string) ([]*gitlab.GroupIteration, error) {
	var (
		itopt gitlab.ListGroupIterationsOptions
		gi    []*gitlab.GroupIteration
		err   error
	)

	itopt.State = gitlab.String("current")
	gi, _, err = c.GL.GroupIterations.ListGroupIterations(gid, &itopt)
	return gi, err
}

func (c *Client) ListGroupIssues(gid string, iteration int) ([]*gitlab.Issue, error) {
	var (
		grpiss gitlab.ListGroupIssuesOptions
		iss    []*gitlab.Issue
		err    error
	)
	grpiss.IterationID = gitlab.Int(iteration)
	iss, _, err = c.GL.Issues.ListGroupIssues(gid, &grpiss)
	return iss, err

}
