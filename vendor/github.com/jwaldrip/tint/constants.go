package tint

const escape = "\x1b"

// DefaultColor for ANSI 16
const DefaultColor = 39

const (
	reset = iota + 0
	bold
	dim
	_ // not used
	underline
	blink
	invert
	hidden
)

const (
	// Black for ANSI 16
	Black = (iota + 30)
	// Red for ANSI 16
	Red
	// Green for ANSI Green
	Green
	// Yellow for ANSI 16
	Yellow
	// Blue for ANSI 16
	Blue
	// Magenta for ANSI 16
	Magenta
	// Cyan for ANSI 16
	Cyan
	// LightGrey for ANSI 16
	LightGrey
)

const (
	// LightBlack for ANSI 16
	LightBlack = (iota + 90)
	// LightRed for ANSI 16
	LightRed
	// LightGreen for ANSI 16
	LightGreen
	// LightYellow for ANSI 16
	LightYellow
	// LightBlue for ANSI 16
	LightBlue
	// LightMagenta for ANSI 16
	LightMagenta
	// LightCyan for ANSI 16
	LightCyan
	// White for ANSI 16
	White
)
