package casparcg

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/overlayfox/casparcg-amcp-go/types"
)

// Client represents a connection to a CasparCG server.
type Client struct {
	host string
	port int
	conn net.Conn
	mu   sync.Mutex
}

// NewClient creates a new CasparCG client.
func NewClient(host string, port int) *Client {
	return &Client{
		host: host,
		port: port,
	}
}

// Connect establishes a connection to the CasparCG server.
func (c *Client) Connect(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var d net.Dialer
	addr := net.JoinHostPort(c.host, strconv.Itoa(c.port))
	conn, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to connect to CasparCG server: %w", err)
	}

	c.conn = conn
	return nil
}

// Close closes the connection to the CasparCG server.
func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// Send sends a command to the CasparCG server and returns the response.
func (c *Client) Send(command interface{ String() string }) ([]string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn == nil {
		return nil, errors.New("not connected to server")
	}

	cmdStr := command.String() + "\r\n"
	_, err := c.conn.Write([]byte(cmdStr))
	if err != nil {
		return nil, fmt.Errorf("failed to send command: %w", err)
	}

	return c.readResponse()
}

// readResponse reads and parses a response from the CasparCG server.
func (c *Client) readResponse() ([]string, error) {
	reader := bufio.NewReader(c.conn)

	// Read the first line to get the response code
	rawFirstLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, types.CasparCGError{
			Code:    0,
			Message: fmt.Sprintf("failed to read response: %v", err),
		}
	}

	firstLine := strings.TrimSpace(rawFirstLine)
	parts := strings.SplitN(firstLine, " ", 2)

	if len(parts) < 1 {
		return nil, types.CasparCGError{
			Code:    0,
			Message: "invalid response format",
		}
	}

	// Try to parse the first part as a numeric code
	code, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, types.CasparCGError{
			Code:    0,
			Message: fmt.Sprintf("could not parse response code: %v", err),
		}
	}

	casparErr := types.CasparCGError{}
	casparErr.Code = code
	if len(parts) > 1 {
		casparErr.Message = parts[1]
	}

	// Almost any response code can be followed by multiline data.
	// Which is why we check for the presence of data for 10 milliseconds after receiving the first line.
	// If no data is received, we assume there is none and return the response.
	if casparErr.Code >= 200 && casparErr.Code < 300 {
		if err := c.conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond)); err != nil {
			return nil, types.CasparCGError{
				Code:    0,
				Message: fmt.Sprintf("failed to set read deadline: %v", err),
			}
		}
		defer func() {
			_ = c.conn.SetReadDeadline(time.Time{})
		}()

		data := make([]string, 0)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}

			// AMCP Terminal Delimiter: A line that is JUST \r\n
			// After TrimSpace, this becomes an empty string.
			if line == "\r\n" || line == "\n" {
				break
			}

			data = append(data, strings.TrimSpace(line))
		}
		return data, nil
	}
	return nil, casparErr
}
