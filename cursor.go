package goansi

import "fmt"

func (as *ANSIString) CursorToStartOfLine() *ANSIString { return as.appendPrefix("1G") }
func (as *ANSIString) CursorHide() *ANSIString          { return as.appendPrefix("?25l") }
func (as *ANSIString) CursorShow() *ANSIString          { return as.appendPrefix("?25h") }
func (as *ANSIString) CursorStore() *ANSIString         { return as.appendPrefix("s") }
func (as *ANSIString) CursorRestore() *ANSIString       { return as.appendPrefix("u") }
func (as *ANSIString) CursorUp(n int) *ANSIString       { return as.appendPrefix(fmt.Sprintf("%dA", n)) }
func (as *ANSIString) CursorDown(n int) *ANSIString     { return as.appendPrefix(fmt.Sprintf("%dB", n)) }
func (as *ANSIString) CursorLeft(n int) *ANSIString     { return as.appendPrefix(fmt.Sprintf("%dD", n)) }
func (as *ANSIString) CursorRight(n int) *ANSIString    { return as.appendPrefix(fmt.Sprintf("%dC", n)) }
func (as *ANSIString) CursorColumn(n int) *ANSIString   { return as.appendPrefix(fmt.Sprintf("%dG", n)) }
func (as *ANSIString) CursorRow(n int) *ANSIString      { return as.appendPrefix(fmt.Sprintf("%d;H", n)) }
func (as *ANSIString) Cursor(x, y int) *ANSIString {
	return as.appendPrefix(fmt.Sprintf("%d;%dH", x, y))
}
