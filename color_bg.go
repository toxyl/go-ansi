package goansi

func (as *ANSIString) BGClear() *ANSIString {
	as.style.color.bg = -1
	return as.Text()
}

// BG sets the background to the given color. If `text` is given the change will only apply to the given text.
func (as *ANSIString) BG(color int, text ...string) *ANSIString {
	if len(text) > 0 {
		colorPrev := as.style.color.bg
		defer func(c int) { as.style.color.bg = c }(colorPrev)
	}
	as.style.color.bg = color
	return as.Text(text...)
}

func (as *ANSIString) DarkBlueBG(str ...string) *ANSIString   { return as.BG(DarkBlue, str...) }
func (as *ANSIString) BlueBG(str ...string) *ANSIString       { return as.BG(Blue, str...) }
func (as *ANSIString) DarkGreenBG(str ...string) *ANSIString  { return as.BG(DarkGreen, str...) }
func (as *ANSIString) LightBlueBG(str ...string) *ANSIString  { return as.BG(LightBlue, str...) }
func (as *ANSIString) OliveGreenBG(str ...string) *ANSIString { return as.BG(OliveGreen, str...) }
func (as *ANSIString) GreenBG(str ...string) *ANSIString      { return as.BG(Green, str...) }
func (as *ANSIString) CyanBG(str ...string) *ANSIString       { return as.BG(Cyan, str...) }
func (as *ANSIString) PurpleBG(str ...string) *ANSIString     { return as.BG(Purple, str...) }
func (as *ANSIString) DarkOrangeBG(str ...string) *ANSIString { return as.BG(DarkOrange, str...) }
func (as *ANSIString) DarkYellowBG(str ...string) *ANSIString { return as.BG(DarkYellow, str...) }
func (as *ANSIString) LimeBG(str ...string) *ANSIString       { return as.BG(Lime, str...) }
func (as *ANSIString) DarkRedBG(str ...string) *ANSIString    { return as.BG(DarkRed, str...) }
func (as *ANSIString) RedBG(str ...string) *ANSIString        { return as.BG(Red, str...) }
func (as *ANSIString) PinkBG(str ...string) *ANSIString       { return as.BG(Pink, str...) }
func (as *ANSIString) OrangeBG(str ...string) *ANSIString     { return as.BG(Orange, str...) }
func (as *ANSIString) YellowBG(str ...string) *ANSIString     { return as.BG(Yellow, str...) }
func (as *ANSIString) DarkGrayBG(str ...string) *ANSIString   { return as.BG(DarkGray, str...) }
func (as *ANSIString) MediumGrayBG(str ...string) *ANSIString { return as.BG(MediumGray, str...) }
func (as *ANSIString) GrayBG(str ...string) *ANSIString       { return as.BG(Gray, str...) }
func (as *ANSIString) WhiteBG(str ...string) *ANSIString      { return as.BG(White, str...) }
func (as *ANSIString) BrightYellowBG(str ...string) *ANSIString {
	return as.BG(BrightYellow, str...)
}
