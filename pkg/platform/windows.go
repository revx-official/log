//go:build windows

package platform

import (
	"os"

	"golang.org/x/sys/windows"
)

// Description:
//
//	Package initialization function.
//	This function is executed only on Windows.
//	Prepares the windows console to accept ASCII escape sequences.
func init() {
	var outMode uint32

	out := windows.Handle(os.Stdout.Fd())
	err := windows.GetConsoleMode(out, &outMode)

	if err != nil {
		println("platform: failed to setup windows console")
		os.Exit(1)
	}

	outMode |= windows.ENABLE_PROCESSED_OUTPUT | windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	_ = windows.SetConsoleMode(out, outMode)
}
