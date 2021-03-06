package etcdtest

import (
	"github.com/coreos/go-etcd/etcd"
)

// Client represents a fake etcd client. Used for testing.
type Client struct {
	Responses map[string][]*etcd.Response
}

// Get mimics the etcd.Client.Get() method.
func (c *Client) Get(key string) ([]*etcd.Response, error) {
	return c.Responses[key], nil
}

// AddResponses adds or updates the Client.Responses map.
func (c *Client) AddResponse(key string, response []*etcd.Response) {
	c.Responses[key] = response
}

// NewClient returns a fake etcd client.
func NewClient() *Client {
	responses := make(map[string][]*etcd.Response)
	return &Client{responses}
}
