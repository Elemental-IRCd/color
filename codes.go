package color

// A color code that mIRC defines
type Color int

// Color codes that mIRC uses. Names taken from http://www.mirc.com/colors.html.
const (
	ColorWhite Color = iota
	ColorBlack
	ColorBlue
	ColorGreen
	ColorLightRed
	ColorBrown
	ColorPurple
	ColorOrange
	ColorYellow
	ColorLightGreen
	ColorCyan
	ColorLightCyan
	ColorLightBlue
	ColorPink
	ColorGray
	ColorLightGray
)

//go:generate stringer -type=Color

// The delimiter for color codes
const (
	ColorCodeDelim = '\x03'
)
