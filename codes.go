package color

// A color code that mIRC defines
type Color int

// Color codes that mIRC uses. Names taken from http://www.mirc.com/colors.html.
const (
	White Color = iota
	Black
	Blue
	Green
	LightRed
	Brown
	Purple
	Orange
	Yellow
	LightGreen
	Cyan
	LightCyan
	LightBlue
	Pink
	Gray
	LightGray
)

//go:generate stringer -type=Color

// The delimiter for color codes
const (
	ColorCodeDelim = '\x03'
)
