package goansi

func (as *ANSIString) FGClear() *ANSIString {
	as.style.color.fg = -1
	return as.Text()
}

// FG sets the foreground to the given color. If `text` is given the change will only apply to the given text.
func (as *ANSIString) FG(color int, text ...string) *ANSIString {
	if len(text) > 0 {
		colorPrev := as.style.color.fg
		defer func(c int) { as.style.color.fg = c }(colorPrev)
	}
	as.style.color.fg = color
	return as.Text(text...)
}

func (as *ANSIString) DarkBlue(str ...string) *ANSIString   { return as.FG(DarkBlue, str...) }
func (as *ANSIString) Blue(str ...string) *ANSIString       { return as.FG(Blue, str...) }
func (as *ANSIString) DarkGreen(str ...string) *ANSIString  { return as.FG(DarkGreen, str...) }
func (as *ANSIString) LightBlue(str ...string) *ANSIString  { return as.FG(LightBlue, str...) }
func (as *ANSIString) OliveGreen(str ...string) *ANSIString { return as.FG(OliveGreen, str...) }
func (as *ANSIString) Green(str ...string) *ANSIString      { return as.FG(Green, str...) }
func (as *ANSIString) Cyan(str ...string) *ANSIString       { return as.FG(Cyan, str...) }
func (as *ANSIString) Purple(str ...string) *ANSIString     { return as.FG(Purple, str...) }
func (as *ANSIString) DarkOrange(str ...string) *ANSIString { return as.FG(DarkOrange, str...) }
func (as *ANSIString) DarkYellow(str ...string) *ANSIString { return as.FG(DarkYellow, str...) }
func (as *ANSIString) Lime(str ...string) *ANSIString       { return as.FG(Lime, str...) }
func (as *ANSIString) DarkRed(str ...string) *ANSIString    { return as.FG(DarkRed, str...) }
func (as *ANSIString) Red(str ...string) *ANSIString        { return as.FG(Red, str...) }
func (as *ANSIString) Pink(str ...string) *ANSIString       { return as.FG(Pink, str...) }
func (as *ANSIString) Orange(str ...string) *ANSIString     { return as.FG(Orange, str...) }
func (as *ANSIString) Yellow(str ...string) *ANSIString     { return as.FG(Yellow, str...) }
func (as *ANSIString) DarkGray(str ...string) *ANSIString   { return as.FG(DarkGray, str...) }
func (as *ANSIString) MediumGray(str ...string) *ANSIString { return as.FG(MediumGray, str...) }
func (as *ANSIString) Gray(str ...string) *ANSIString       { return as.FG(Gray, str...) }
func (as *ANSIString) White(str ...string) *ANSIString      { return as.FG(White, str...) }
func (as *ANSIString) BrightYellow(str ...string) *ANSIString {
	return as.FG(BrightYellow, str...)
}
