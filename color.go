package goansi

func (as *ANSIString) Color(colorBG, colorFG int, text ...string) *ANSIString {
	if len(text) > 0 {
		colorBGPrev := as.style.color.bg
		colorFGPrev := as.style.color.fg
		defer func(cbg, cfg int) {
			as.style.color.bg = cbg
			as.style.color.fg = cfg
		}(colorBGPrev, colorFGPrev)
	}
	as.style.color.bg = colorBG
	as.style.color.fg = colorFG
	return as.Text(text...)
}
