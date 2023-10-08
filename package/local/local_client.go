package local

import (
	"github.com/d-strobel/gowindows/package/connection"
)

type Client struct {
	Connection *connection.Connection
}

// NewClient returns a Client for the local package
func NewClient(conn *connection.Connection) *Client {
	return &Client{Connection: conn}
}
