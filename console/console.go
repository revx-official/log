package console

import (
	"fmt"
)

// Description:
//
//	Writes a formatted message to the console.
//
// Parameters:
//
//	format 	The string used to format the given arguments.
//	args 	The arguments to write.
func WriteFormat(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// Description:
//
//	Writes a message to the console, terminated with a new line symbol.
//
// Parameters:
//
//	args 	The arguments to write.
func WriteLine(args ...interface{}) {
	fmt.Println(args...)
}
