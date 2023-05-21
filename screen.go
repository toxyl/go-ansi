package goansi

import "fmt"

func (as *ANSIString) ScrollUp(n int) *ANSIString   { return as.appendPrefix(fmt.Sprintf("%dS", n)) }
func (as *ANSIString) ScrollDown(n int) *ANSIString { return as.appendPrefix(fmt.Sprintf("%dT", n)) }
