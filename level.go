package log

// Description:
//
//	The log level type definition.
//	Log levels are represented by numbers.
type LogLevel = uint32

const (
	// The trace log level.
	LogLevelTrace LogLevel = iota

	// The debug log level.
	LogLevelDebug LogLevel = iota

	// The info log level.
	LogLevelInfo LogLevel = iota

	// The warn log level.
	LogLevelWarn LogLevel = iota

	// The error log level.
	LogLevelError LogLevel = iota

	// The fatal log level.
	LogLevelFatal LogLevel = iota
)
