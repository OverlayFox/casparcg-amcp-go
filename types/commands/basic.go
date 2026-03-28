package commands

import (
	"github.com/overlayfox/casparcg-amcp-go/types"
)

type DirectCommandLogLevel struct {
	Level types.LogLevel
}

func (c DirectCommandLogLevel) String() string {
	return "LOG LEVEL " + string(c.Level)
}

type DirectCommandPing struct {
	Token *string
}

func (c DirectCommandPing) String() string {
	cmd := "PING"
	return appendString(cmd, c.Token)
}

type CommandBye struct{}

func (c CommandBye) String() string {
	return "BYE"
}

type CommandKill struct{}

func (c CommandKill) String() string {
	return "KILL"
}

type CommandRestart struct{}

func (c CommandRestart) String() string {
	return "RESTART"
}
