package goansi

import "fmt"

func (as *ANSIString) MoveHome() *ANSIString {
	as.posX = 0
	as.posY = 0
	return as.appendPrefix("H")
}

func (as *ANSIString) MoveTo(x, y int) *ANSIString {
	as.posX = x
	as.posY = y
	if x < 0 {
		as.posX = as.termWidth + x
	}
	if y < 0 {
		as.posY = as.termHeight + y + 1
	}
	return as.appendPrefix(fmt.Sprintf("%d;%dH", as.posY, as.posX))
}

func (as *ANSIString) MoveToRow(row int) *ANSIString {
	return as.MoveTo(as.posX, row)
}

func (as *ANSIString) MoveToFirstRow() *ANSIString {
	return as.MoveToRow(0)
}

func (as *ANSIString) MoveToLastRow() *ANSIString {
	return as.MoveToRow(-1)
}

func (as *ANSIString) MoveToColumn(column int) *ANSIString {
	return as.MoveTo(column, as.posY)
}

func (as *ANSIString) MoveToFirstColumn() *ANSIString {
	return as.MoveToColumn(0)
}

func (as *ANSIString) MoveToLastColumn() *ANSIString {
	return as.MoveToColumn(-1)
}
