package goansi

import (
	"strings"

	"github.com/toxyl/go-ansi/utils"
)

func (as *ANSIString) Fill(color, x, y, w, h int) *ANSIString {
	x, y, w, h = utils.TermWindow.SelectArea(x, y, w, h)

	text := strings.Repeat(" ", w) + "\n"
	for i := 0; i < h; i++ {
		as.MoveTo(x, y+i).BG(color).Text(text)
	}
	return as
}
