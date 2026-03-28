package casparcg

import (
	"github.com/overlayfox/casparcg-amcp-go/types"
	"github.com/overlayfox/casparcg-amcp-go/types/commands"
)

// sendCommand abstracts sending a command that does not expect a response value.
func (b *Client) sendCommand(cmd interface{ String() string }) error {
	_, err := b.Send(cmd)
	return err
}

// LogLevel sets the log level.
func (c *Client) LogLevel(level types.LogLevel) error {
	cmd := commands.DirectCommandLogLevel{
		Level: level,
	}
	return c.sendCommand(cmd)
}

// Ping sends a ping command.
func (c *Client) Ping(token *string) (string, error) {
	cmd := commands.DirectCommandPing{
		Token: token,
	}
	data, err := c.Send(cmd)
	if err != nil {
		return "", err
	}
	return data[0], nil
}

// Bye closes the connection.
func (c *Client) Bye() error {
	cmd := commands.CommandBye{}
	_, err := c.Send(cmd)
	return err
}

// Kill kills the server.
func (c *Client) Kill() error {
	cmd := commands.CommandKill{}
	_, err := c.Send(cmd)
	return err
}

// Restart restarts the server.
func (c *Client) Restart() error {
	cmd := commands.CommandRestart{}
	_, err := c.Send(cmd)
	return err
}
