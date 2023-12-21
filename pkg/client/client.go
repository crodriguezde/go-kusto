package client

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/crodriguezde/go-kusto/pkg/conn"
)

type Client struct {
	cred     azcore.TokenCredential
	http     *http.Client
	endpoint string
	conn     *conn.Conn
}

func New(options ...ClientOption) (*Client, error) {
	var err error

	c := &Client{}

	for _, option := range options {
		option(c)
	}

	if c.http == nil {
		c.http = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}

	conn, err := conn.NewConn(c.endpoint, c.cred, c.http)
	if err != nil {
		return nil, err
	}

	c.conn = conn

	return c, nil
}

func (c *Client) Query(ctx context.Context) error {
	return c.conn.Query(ctx, "eventmapper", "logs_eventmapper_v2 | take 1\n", nil)
}
