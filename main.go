package goansi

import (
	"fmt"
	"strings"
)

type ANSIString struct {
	tokens     []string
	style      *style
	termWidth  int
	termHeight int
	posX       int
	posY       int
}

func (as *ANSIString) Clear() *ANSIString {
	as.style.clear()
	return as.Text()
}

func (as *ANSIString) String() string {
	return strings.Join(as.tokens, "")
}

func (as *ANSIString) Print() {
	fmt.Print(as.String())
}

func (as *ANSIString) Println() {
	fmt.Print(as.Reset().String() + "\n")
}

func (as *ANSIString) add(str ...string) *ANSIString {
	as.tokens = append(as.tokens, str...)
	return as
}

func (as *ANSIString) appendPrefix(str string) *ANSIString {
	return as.add("\033[" + str)
}

func New() *ANSIString {
	w, h := getTermSize()
	as := &ANSIString{
		tokens: []string{},
		style: &style{
			color: struct {
				fg int
				bg int
			}{
				fg: -1,
				bg: -1,
			},
			bold:          false,
			italic:        false,
			underline:     false,
			strikeThrough: false,
		},
		termWidth:  w,
		termHeight: h,
	}
	return as
}
