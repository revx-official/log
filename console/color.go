package console

import (
	"fmt"
	"strconv"
	"strings"
)

// Description:
//
//	Type definition for console color or emphasis attributes.
//	Those attributes are used within escape sequences.
type ConsoleAttribute = uint32

const (
	// The console escape character.
	ConsoleEscape = "\x1b"
)

const (
	// Resets all current console attributes.
	Reset ConsoleAttribute = iota

	// Bold text emphasis.
	EmphasisBold

	// Faint text emphasis.
	EmphasisFaint

	// Italic text emphasis.
	EmphasisItalic

	// Underline text emphasis.
	EmphasisUnderline

	// Blink text emphasis (slow).
	EmphasisBlinkSlow

	// Blink text emphasis (fast).
	EmphasisBlinkFast

	// Reverse video text emphasis.
	EmphasisReverseVideo

	// Concealed text emphasis.
	EmphasisConcealed

	// Crossed out text emphasis.
	EmphasisCrossedOut
)

const (
	// Escape code for console foreground color: black.
	FgColorBlack ConsoleAttribute = iota + 30

	// Escape code for console foreground color: red.
	FgColorRed

	// Escape code for console foreground color: green.
	FgColorGreen

	// Escape code for console foreground color: yellow.
	FgColorYellow

	// Escape code for console foreground color: blue.
	FgColorBlue

	// Escape code for console foreground color: magenta.
	FgColorMagenta

	// Escape code for console foreground color: cyan.
	FgColorCyan

	// Escape code for console foreground color: white.
	FgColorWhite
)

const (
	// Escape code for high intensity console foreground color: black.
	FgColorIntenseBlack ConsoleAttribute = iota + 90

	// Escape code for high intensity console foreground color: red.
	FgColorIntenseRed

	// Escape code for high intensity console foreground color: green.
	FgColorIntenseGreen

	// Escape code for high intensity console foreground color: yellow.
	FgColorIntenseYellow

	// Escape code for high intensity console foreground color: blue.
	FgColorIntenseBlue

	// Escape code for high intensity console foreground color: magenta.
	FgColorIntenseMagenta

	// Escape code for high intensity console foreground color: cyan.
	FgColorIntenseCyan

	// Escape code for high intensity console foreground color: white.
	FgColorIntenseWhite
)

const (
	// Escape code for console background color: black.
	BgColorBlack ConsoleAttribute = iota + 40

	// Escape code for console background color: red.
	BgColorRed

	// Escape code for console background color: green.
	BgColorGreen

	// Escape code for console background color: yellow.
	BgColorYellow

	// Escape code for console background color: blue.
	BgColorBlue

	// Escape code for console background color: magenta.
	BgColorMagenta

	// Escape code for console background color: cyan.
	BgColorCyan

	// Escape code for console background color: white.
	BgColorWhite
)

const (
	// Escape code for high intensity console background color: black.
	BgColorIntenseBlack ConsoleAttribute = iota + 100

	// Escape code for high intensity console background color: red.
	BgColorIntenseRed

	// Escape code for high intensity console background color: green.
	BgColorIntenseGreen

	// Escape code for high intensity console background color: yellow.
	BgColorIntenseYellow

	// Escape code for high intensity console background color: blue.
	BgColorIntenseBlue

	// Escape code for high intensity console background color: magenta.
	BgColorIntenseMagenta

	// Escape code for high intensity console background color: cyan.
	BgColorIntenseCyan

	// Escape code for high intensity console background color: white.
	BgColorIntenseWhite
)

// Description:
//
//	Represents a console color.
//	Every console color can have multiple attributes, e.g. a color and a text emphasis attribute.
type ConsoleColor struct {
	Attributes []ConsoleAttribute
}

// Description:
//
//	Creates a new console color.
//
// Parameters:
//
//	attributes The initial attributes of the console color.
//
// Returns:
//
//	The created console color.
func NewConsoleColor(attributes ...ConsoleAttribute) *ConsoleColor {
	return &ConsoleColor{
		Attributes: attributes,
	}
}

// Description:
//
//	Wraps a formatted string with the escape sequences describing this color.
//	Works exactly like a the 'Sprintf' function.
//	Includes a reset sequence.
//
// Parameters:
//
//	format 	The format string.
//	args	The arguments used by the format string.
//
// Returns:
//
//	The formatted message including all ANSI escape codes describing this color.
func (color *ConsoleColor) Wrap(format string, args ...interface{}) string {
	sequence := color.EscapeSequence()
	message := fmt.Sprintf(format, args...)

	return fmt.Sprintf(
		"%s[%sm%s%s[%dm",
		ConsoleEscape,
		sequence,
		message,
		ConsoleEscape,
		Reset,
	)
}

// Description:
//
//	Returns the 'Wrap' function of the 'ConsoleColor' type.
//	The returned function can be used as a text wrapper.
//
// Example:
//
//	color := console.NewConsoleColor(console.FgColorGreen)
//	green := color.WrapFunc()
//	fmt.Println(green("%s", "Hello there!"))
//
// Returns:
//
//	The 'Wrap' function of the 'ConsoleColor' type.
func (color *ConsoleColor) WrapFunc() func(format string, args ...interface{}) string {
	return color.Wrap
}

// Description:
//
//	Wraps a formatted string with the escape sequences describing this color and directly prints the result.
//	Works exactly like a the 'Printf' function.
//	Includes a reset sequence.
//
// Parameters:
//
//	format 	The format string.
//	args	The arguments used by the format string.
func (color *ConsoleColor) Printf(format string, args ...interface{}) {
	sequence := color.EscapeSequence()
	message := fmt.Sprintf(format, args...)

	fmt.Printf(
		"%s[%sm%s%s[%dm",
		ConsoleEscape,
		sequence,
		message,
		ConsoleEscape,
		Reset,
	)
}

// Description:
//
//	Returns the 'Printf' function of the 'ConsoleColor' type.
//	The returned function can be used as a printer.
//
// Example:
//
//	color := console.NewConsoleColor(console.FgColorGreen)
//	green := color.PrintFunc()
//	green("%s", "Hello there!")
//
// Returns:
//
//	The 'Printf' function of the 'ConsoleColor' type.
func (color *ConsoleColor) PrintFunc() func(format string, args ...interface{}) {
	return color.Printf
}

// Description:
//
//	Creates the string escape sequence for this color.
//	The sequence does not include the '[' and 'm' formatting characters.
//
// Returns:
//
//	The escape sequence for this color.
func (color *ConsoleColor) EscapeSequence() string {
	count := len(color.Attributes)
	format := make([]string, count)

	for index, attribute := range color.Attributes {
		format[index] = strconv.Itoa(int(attribute))
	}

	return strings.Join(format, ";")
}
