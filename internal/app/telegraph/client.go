package telegraph

import (
	"fmt"
	"github.com/TechMinerApps/telegraph"
)

const (
	authorName = "Referral Bot"
)

type Client struct {
	client  telegraph.Client
	account telegraph.Account
}

func New(name, token string) (*Client, error) {
	client, err := telegraph.NewClientWithToken(token)
	if err != nil {
		return nil, fmt.Errorf("cannot init telegraph service: %v", err)
	}
	return &Client{
		client: client,
	}, nil
}

func (c *Client) CreatePage(title, content string) (string, error) {
	page, err := c.client.CreatePage(telegraph.Page{
		Title:      title,
		AuthorName: authorName,
		Content:    []telegraph.Node{content},
		Views:      0,
		CanEdit:    false,
	}, true)
	if err != nil {
		return "", err
	}

	return page.URL, nil
}
