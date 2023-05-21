package goansi

import (
	"strings"

	"github.com/toxyl/go-ansi/utils"
)

func (as *ANSIString) Fill(color, x, y, w, h int) *ANSIString {
	x, y, _, y2 := utils.Abs(x, y, w, h)
	text := strings.Repeat(" ", w)
	for y < y2 {
		as.MoveTo(x, y).BG(color, text)
		y++
	}
	return as
}
