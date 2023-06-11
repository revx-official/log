package log

import (
	"github.com/revx-official/log/pkg/console"
)

// Description:
//
//	The global log level.
var Level = core.LogLevelInfo

// Description:
//
//	Traces a formatted message.
//
// Parameters:
//
//	format 	The string used to format the given arguments.
//	args	The arguments to log.
func Tracef(format string, args ...interface{}) {
	if Level > LogLevelTrace {
		return
	}

	console.Tracef(format, args...)
}

// Description:
//
//	Logs a formatted debug message.
//
// Parameters:
//
//	format 	The string used to format the given arguments.
//	args	The arguments to log.
func Debugf(format string, args ...interface{}) {
	if Level > LogLevelDebug {
		return
	}

	console.Debugf(format, args...)
}

// Description:
//
//	Logs a formatted info message.
//
// Parameters:
//
//	format 	The string used to format the given arguments.
//	args	The arguments to log.
func Infof(format string, args ...interface{}) {
	if Level > LogLevelInfo {
		return
	}

	console.Infof(format, args...)
}

// Description:
//
//	Logs a formatted warning message.
//
// Parameters:
//
//	format 	The string used to format the given arguments.
//	args	The arguments to log.
func Warnf(format string, args ...interface{}) {
	if Level > LogLevelWarn {
		return
	}

	console.Warnf(format, args...)
}

// Description:
//
//	Logs a formatted error message.
//
// Parameters:
//
//	format 	The string used to format the given arguments.
//	args	The arguments to log.
func Errorf(format string, args ...interface{}) {
	if Level > LogLevelError {
		return
	}

	console.Errorf(format, args...)
}

// Description:
//
//	Logs a formatted fatal message.
//
// Parameters:
//
//	format 	The string used to format the given arguments.
//	args	The arguments to log.
func Fatalf(format string, args ...interface{}) {
	console.Fatalf(format, args...)
}
