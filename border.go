package goansi

import "strings"

func (as *ANSIString) Box(x, y, w, h int, label string) *ANSIString {
	y++
	h--
	top := "┌" + strings.Repeat("─", w-2) + "┐\n"
	bottom := "└" + strings.Repeat("─", w-2) + "┘\n"
	as.MoveTo(x, y).Text(top)
	for i := 1; i < h; i++ {
		x2 := x + w
		if x > 0 {
			x2--
		}
		as.MoveToRow(y + i)
		as.MoveToColumn(x).Text("│")
		as.MoveToColumn(x2).Text("│")
	}
	as.MoveTo(x, y+h).Text(bottom)
	if label != "" {
		as.MoveTo(x+3, y).Bold(" " + label + " ")
	}

	return as
}
