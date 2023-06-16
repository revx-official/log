package log

// Description:
//
//	The log level type definition.
//	Log levels are represented by numbers.
type LogLevel = uint32

const (
	// The trace log level.
	LevelTrace LogLevel = iota

	// The debug log level.
	LevelDebug LogLevel = iota

	// The info log level.
	LevelInfo LogLevel = iota

	// The warn log level.
	LevelWarn LogLevel = iota

	// The error log level.
	LevelError LogLevel = iota

	// The fatal log level.
	LevelFatal LogLevel = iota
)
