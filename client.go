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
	user, _, err := c.GL.Users.CurrentUser()
	if err != nil {
		return nil, err
	}

	glUser := &GitlabUser{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Profile:  user.WebURL,
	}

	return glUser, nil
}
