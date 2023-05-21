package goansi

func (as *ANSIString) Text(text ...string) *ANSIString {
	return as.add(as.style.render(text...))
}

func (as *ANSIString) Bold(text ...string) *ANSIString {
	defer func() { as.style.bold = false }()
	as.style.bold = true
	return as.Text(text...)
}

func (as *ANSIString) Italic(text ...string) *ANSIString {
	defer func() { as.style.italic = false }()
	as.style.italic = true
	return as.Text(text...)
}

func (as *ANSIString) Underline(text ...string) *ANSIString {
	defer func() { as.style.underline = false }()
	as.style.underline = true
	return as.Text(text...)
}

func (as *ANSIString) StrikeThrough(text ...string) *ANSIString {
	defer func() { as.style.strikeThrough = false }()
	as.style.strikeThrough = true
	return as.Text(text...)
}

func (as *ANSIString) Hide(text ...string) *ANSIString {
	defer func() { as.style.textHidden = false }()
	as.style.textHidden = true
	return as.Text(text...)
}

func (as *ANSIString) Reset() *ANSIString { return as.appendPrefix("0m") }
