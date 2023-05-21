package goansi

import (
	"strings"

	"github.com/toxyl/go-ansi/utils"
)

func (as *ANSIString) BoxCustom(x, y, w, h int, label, topLeft, top, topRight, left, center, right, bottomLeft, bottom, bottomRight string) *ANSIString {
	x, y, _, y2 := utils.Abs(x, y, w, h)
	top = topLeft + strings.Repeat(top, w-2) + topRight + "\n"
	bottom = bottomLeft + strings.Repeat(bottom, w-2) + bottomRight + "\n"
	as.MoveTo(x, y).Text(top)
	y++
	center = left + strings.Repeat(center, w-2) + right
	for y < y2 {
		as.MoveTo(x, y).Text(center)
		y++
	}
	as.MoveTo(x, y2).Text(bottom)
	if label != "" {
		as.MoveTo(x+1, y-h).Bold(" " + label + " ")
	}

	return as
}

func (as *ANSIString) Box(x, y, w, h int, label string) *ANSIString {
	x, y, x2, y2 := utils.Abs(x, y, w, h)
	top := "┌" + strings.Repeat("─", w-2) + "┐\n"
	bottom := "└" + strings.Repeat("─", w-2) + "┘\n"
	as.MoveTo(x, y).Text(top)
	y++
	for y < y2 {
		as.MoveTo(x, y).Text("│")
		as.MoveTo(x2, y).Text("│")
		y++
	}
	as.MoveTo(x, y2).Text(bottom)
	if label != "" {
		as.MoveTo(x+1, y-h).Bold(" " + label + " ")
	}

	return as
}
