package color

import "fmt"

// ColoredText is a section of text that is colored with mIRC color codes.
type ColoredText struct {
	Foreground Color  `json:"foreground"`
	Background Color  `json:"background"`
	Content    string `json:"content"`
}

func (c ColoredText) String() string {
	if c.Background != None {
		return fmt.Sprintf("%c%d,%d%s%c", ColorCodeDelim, c.Foreground, c.Background, c.Content, ColorCodeDelim)
	}

	return fmt.Sprintf("%c%d%s%c", ColorCodeDelim, c.Foreground, c.Content, ColorCodeDelim)
}
