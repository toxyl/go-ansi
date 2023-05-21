package goansi

import (
	"fmt"
	"strings"
)

type style struct {
	color struct {
		fg int
		bg int
	}
	bold          bool
	italic        bool
	underline     bool
	strikeThrough bool
	textHidden    bool
}

func (s *style) String() string {
	prefix := "\033["
	res := prefix + "0m"
	if s.bold {
		res += prefix + "1m"
	}
	if s.italic {
		res += prefix + "3m"
	}
	if s.underline {
		res += prefix + "4m"
	}
	if s.strikeThrough {
		res += prefix + "9m"
	}
	if s.textHidden {
		res += prefix + "8m"
	}
	if s.color.fg >= 0 {
		res += prefix + fmt.Sprintf("38;5;%dm", MapColor(s.color.fg))
	}
	if s.color.bg >= 0 {
		res += prefix + fmt.Sprintf("48;5;%dm", MapColor(s.color.bg))
	}
	return res
}

func (s *style) render(str ...string) string {
	return fmt.Sprintf("%s%s", s.String(), strings.Join(str, ""))
}

func (s *style) clear() {
	s.color.fg = -1
	s.color.bg = -1
	s.bold = false
	s.italic = false
	s.strikeThrough = false
	s.textHidden = false
	s.underline = false
}
