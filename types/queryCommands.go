package types

type QueryCommandInterface interface {
}

// QueryCommandCINF returns information about a media file.
type QueryCommandCINF struct {
	filename string
}

// QueryCommandCLS lists media files in the media folder.
// Use the command INFO PATHS to get the path to the media folder.
type QueryCommandCLS struct {
	directory *string
}

// QueryCommandFLS lists all fonts in the fonts folder.
// Use the command INFO PATHS to get the path to the fonts folder.
type QueryCommandFLS struct {
}

// QueryCommandTLS lists template files in the templates folder.
// Use the command INFO PATHS to get the path to the templates folder.
type QueryCommandTLS struct {
	directory *string
}

// QueryCommandVersion returns the version of specified component.
type QueryCommandVersion struct {
	component *string
}

type InfoComponent string

const (
	InfoComponentConfig  InfoComponent = "CONFIG"
	InfoComponentPaths   InfoComponent = "PATHS"
	InfoComponentSystem  InfoComponent = "SYSTEM"
	InfoComponentServer  InfoComponent = "SERVER"
	InfoComponentQueues  InfoComponent = "QUEUES"
	InfoComponentThreads InfoComponent = "THREADS"
)

// QueryCommandInfo retrieves a list of available channels
type QueryCommandInfo struct {
	component *InfoComponent
}

// QueryCommandInfoChannel get information about a channel or a specific layer on a channel.
// If layer is ommitted information about the whole channel is returned.
type QueryCommandInfoChannel struct {
	videoChannel int
	layer        *int
}

// QueryCommandInfoTemplate gets information about the specified template.
type QueryCommandInfoTemplate struct {
	template string
}

type QueryCommandInfoDelay struct {
	videoChannel int
	layer        *int
}
