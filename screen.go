package goansi

import (
	"fmt"

	"github.com/toxyl/go-ansi/utils"
)

func (as *ANSIString) ScrollUp(n int) *ANSIString   { return as.appendPrefix(fmt.Sprintf("%dS", n)) }
func (as *ANSIString) ScrollDown(n int) *ANSIString { return as.appendPrefix(fmt.Sprintf("%dT", n)) }

func W() int {
	return utils.TermWindow.W()
}

func H() int {
	return utils.TermWindow.W()
}
