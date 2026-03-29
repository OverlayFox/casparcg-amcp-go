package casparcg

import (
	"strings"

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
	resp, err := c.Send(cmd)
	if err != nil {
		return "", err
	}

	return strings.Join(resp, " "), nil
}

// Diag opens the diagnostic window.
func (c *Client) Diag() error {
	cmd := commands.QueryDiag{}
	return c.sendCommand(cmd)
}

// GLGC releases all the pooled OpenGL resources.
//
// ⚠️ WARNING: May cause a pause on all video channels.
func (b *QueryBuilder) GLGC() error {
	cmd := commands.QueryGLGC{}
	return b.sendCommand(cmd)
}

// Bye closes the connection.
func (c *Client) Bye() error {
	cmd := commands.CommandBye{}
	return c.sendCommand(cmd)
}

// Kill kills the server.
func (c *Client) Kill() error {
	cmd := commands.CommandKill{}
	return c.sendCommand(cmd)
}

// Restart restarts the server.
func (c *Client) Restart() error {
	cmd := commands.CommandRestart{}
	return c.sendCommand(cmd)
}

type ClientHelpCommand struct {
	client *Client
}

// Help shows online help for a specific command or a list of all commands.
func (c *Client) Help() *ClientHelpCommand {
	return &ClientHelpCommand{
		client: c,
	}
}

// Generic shows a detailed description of the specified command, or a list of all commands if no command is specified.
func (b *ClientHelpCommand) Generic(command *string) ([]string, error) {
	cmd := commands.QueryHelp{
		Command: command,
	}
	return b.client.Send(cmd)
}

// Producer shows a detailed description of the specified producer, or a list of all producers if no producer is specified.
func (b *ClientHelpCommand) Producer(producer *string) ([]string, error) {
	cmd := commands.QueryHelpProducer{
		Producer: producer,
	}
	return b.client.Send(cmd)
}

// Consumer shows a detailed description of the specified consumer, or a list of all consumers if no consumer is specified.
func (b *ClientHelpCommand) Consumer(consumer *string) ([]string, error) {
	cmd := commands.QueryHelpConsumer{
		Consumer: consumer,
	}
	return b.client.Send(cmd)
}
