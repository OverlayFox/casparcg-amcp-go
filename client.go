package casparcg

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/overlayfox/casparcg-amcp-go/types/returns"
)

// Client represents a connection to a CasparCG server
type Client struct {
	host string
	port int
	conn net.Conn
	mu   sync.Mutex
}

// Response represents a response from the CasparCG server
type Response struct {
	Code    returns.ReturnCode
	Message string
	Data    []string
}

// NewClient creates a new CasparCG client
func NewClient(host string, port int) *Client {
	return &Client{
		host: host,
		port: port,
	}
}

// Connect establishes a connection to the CasparCG server
func (c *Client) Connect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	addr := net.JoinHostPort(c.host, strconv.Itoa(c.port))
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to connect to CasparCG server: %w", err)
	}

	c.conn = conn
	return nil
}

// Close closes the connection to the CasparCG server
func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// Send sends a command to the CasparCG server and returns the response
func (c *Client) Send(command interface{ String() string }) (*Response, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn == nil {
		return nil, fmt.Errorf("not connected to server")
	}

	cmdStr := command.String() + "\r\n"
	_, err := c.conn.Write([]byte(cmdStr))
	if err != nil {
		return nil, fmt.Errorf("failed to send command: %w", err)
	}

	return c.readResponse()
}

// readResponse reads and parses a response from the CasparCG server
func (c *Client) readResponse() (*Response, error) {
	reader := bufio.NewReader(c.conn)

	// Read the first line to get the response code
	rawFirstLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	firstLine := strings.TrimSpace(rawFirstLine)
	parts := strings.SplitN(firstLine, " ", 2)

	if len(parts) < 1 {
		return nil, fmt.Errorf("invalid response format")
	}

	response := &Response{
		Data: []string{},
	}

	// Try to parse the first part as a numeric code
	code, err := strconv.Atoi(parts[0])
	if err != nil {
		response.Code = returns.ReturnCode(0)
		response.Message = firstLine
		return response, nil
	}
	response.Code = returns.ReturnCode(code)
	if len(parts) > 1 {
		response.Message = parts[1]
	}

	// Almost any response code can be followed by multiline data.
	// Which is why we check for the presence of data for 10 milliseconds after receiving the first line.
	// If no data is received, we assume there is none and return the response.
	if response.Code >= 200 && response.Code < 300 {
		if err := c.conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond)); err != nil {
			return nil, fmt.Errorf("failed to set read deadline: %w", err)
		}
		defer func() {
			_ = c.conn.SetReadDeadline(time.Time{})
		}()
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

			response.Data = append(response.Data, strings.TrimSpace(line))
		}
	}

	return response, nil
}
