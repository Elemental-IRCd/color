// generated by stringer -type=Color; DO NOT EDIT

package color

import "fmt"

const _Color_name = "WhiteBlackBlueGreenLightRedBrownPurpleOrangeYellowLightGreenCyanLightCyanLightBluePinkGrayLightGray"

var _Color_index = [...]uint8{0, 5, 10, 14, 19, 27, 32, 38, 44, 50, 60, 64, 73, 82, 86, 90, 99}

func (i Color) String() string {
	if i < 0 || i+1 >= Color(len(_Color_index)) {
		return fmt.Sprintf("Color(%d)", i)
	}
	return _Color_name[_Color_index[i]:_Color_index[i+1]]
}