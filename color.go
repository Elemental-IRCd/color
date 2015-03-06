package color

// A color code that mIRC defines
type Color int

//go:generate stringer -type=Color

func (c Color) Validate() bool {
	return c >= None && c < (LightGray+1)
}
