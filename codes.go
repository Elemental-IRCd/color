package color

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
	None Color = -1 // No color
)

// The delimiter for color codes
const (
	ColorCodeDelim = '\x03'
)
