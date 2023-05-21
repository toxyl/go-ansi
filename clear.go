package goansi

func (as *ANSIString) ClearCursorToScreenEnd() *ANSIString   { return as.appendPrefix("0J") }
func (as *ANSIString) ClearCursorToScreenStart() *ANSIString { return as.appendPrefix("1J") }
func (as *ANSIString) ClearScreen() *ANSIString              { return as.appendPrefix("2J") }
func (as *ANSIString) ClearLine() *ANSIString                { return as.appendPrefix("2K") }
func (as *ANSIString) ClearToEOL() *ANSIString               { return as.appendPrefix("K") }
func (as *ANSIString) ClearToStart() *ANSIString             { return as.appendPrefix("1K") }
