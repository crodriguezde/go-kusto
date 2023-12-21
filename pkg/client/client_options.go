package client

import (
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

type ClientOption func(c *Client)

func WithTokenCredential(cred azcore.TokenCredential) ClientOption {
	return func(c *Client) {
		c.cred = cred
	}
}

func WithEndpoint(e string) ClientOption {
	return func(c *Client) {
		c.endpoint = e
	}
}

func WithHttp(http *http.Client) ClientOption {
	return func(c *Client) {
		c.http = http
	}
}
