// Package local provides a Go library for handling local Windows functions.
package local

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/d-strobel/gowindows/connection"
	"github.com/d-strobel/gowindows/parsing"
)

// localType is a type constraint for the localRun function, ensuring it works with specific types.
type localType interface {
	Group | []Group | User | []User | GroupMember | []GroupMember
}

// LocalClient represents a client for handling local Windows functions.
type LocalClient struct {
	// Connection represents a connection.Connection object.
	Connection connection.Connection

	// decodeCliXmlErr represents a function that decodes a CLIXML error and returns aa  human readable string.
	decodeCliXmlErr func(string) (string, error)
}

// NewClient returns a new instance of the LocalClient.
func NewClient(conn connection.Connection) *LocalClient {
	return NewClientWithParser(conn, parsing.DecodeCliXmlErr)
}

// NewClientWithParser returns a new instance of the LocalClient.
// It requires a connection and parsing as input parameters.
func NewClientWithParser(conn connection.Connection, parsing func(string) (string, error)) *LocalClient {
	return &LocalClient{Connection: conn, decodeCliXmlErr: parsing}
}

// SID represents the Security Identifier (SID) of a security principal.
// The Value field contains the actual SID value.
type SID struct {
	Value string `json:"Value"`
}

// localRun runs a PowerShell command against a Windows system, handles the command results,
// and unmarshals the output into a local object type.
func localRun[T localType](ctx context.Context, c *LocalClient, cmd string, l *T) error {
	// Run the command
	result, err := c.Connection.RunWithPowershell(ctx, cmd)
	if err != nil {
		return err
	}

	// Handle stderr
	if result.StdErr != "" {
		stderr, err := c.decodeCliXmlErr(result.StdErr)
		if err != nil {
			return err
		}

		return fmt.Errorf("Command:\n%s\n\nPowershell error:\n%s", cmd, stderr)
	}

	if result.StdOut == "" {
		return nil
	}

	// Unmarshal stdout
	if err = json.Unmarshal([]byte(result.StdOut), &l); err != nil {
		return err
	}

	return nil
}
