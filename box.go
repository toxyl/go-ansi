package goansi

import (
	"math"
	"strings"

	"github.com/toxyl/go-ansi/utils"
)

func (as *ANSIString) BoxCustom(x, y, w, h int, label, topLeft, top, topRight, left, center, right, bottomLeft, bottom, bottomRight string) *ANSIString {
	if label != "" {
		label = " " + label + " "
	}

	x, y, w, h = utils.TermWindow.SelectArea(x, y, w, h)

	as.MoveTo(x, y).Text(topLeft).Text(top).Bold(label).Text(strings.Repeat(top, int(math.Max(float64(w-3-len(label)), 0.0)))).Text(topRight).Ln()
	y++
	for i := 0; i < h-2; i++ {
		as.MoveTo(x, y).Text(left).Text(strings.Repeat(center, w-2)).Text(right).Ln()
		y++
	}
	as.MoveTo(x, y).Text(bottomLeft).Text(strings.Repeat(bottom, w-2)).Text(bottomRight)

	return as
}

func (as *ANSIString) Box(x, y, w, h int, label string) *ANSIString {
	return as.BoxCustom(x, y, w, h, label,
		"┌", "─", "┐",
		"│", " ", "│",
		"└", "─", "┘",
	)
}

func (as *ANSIString) Input(x, y, w int, label, text string) *ANSIString {
	h := 3
	x, y, w, h = utils.TermWindow.SelectArea(x, y, w, h)
	as.Box(x, y, w, h, label)

	x += 2
	y++
	w -= 4
	h -= 2

	pad := ""

	if len(text) > w {
		text = "..." + text[len(text)-w+3:] // add ellipsis if text is too long
	} else {
		pad = strings.Repeat(" ", w-len(text))
	}

	as.MoveTo(x, y).Text(strings.TrimSpace(text+pad)).MoveTo(x+len(text), y)

	return as
}

func (as *ANSIString) LogBox(x, y, w, h int, label string, entries []string, wordWrap bool) *ANSIString {
	x, y, w, h = utils.TermWindow.SelectArea(x, y, w, h)
	as.Box(x, y, w, h, label)

	x++
	y++
	w -= 2
	h -= 2

	if len(entries) > h {
		entries = entries[len(entries)-h:] // pick the latest entries if there are more than we can display
	}

	logs := []string{}
	for _, log := range entries {
		lines := []string{}
		if wordWrap && len(log) > w {
			line := ""
			tokens := strings.Split(log, " ")
			for _, t := range tokens {
				if len(line)+len(t)+1 < w {
					line += t + " "
					continue
				}
				lines = append(lines, line)
				line = t + " "
			}
			if line != "" {
				lines = append(lines, line)
				line = ""
			}
		} else {
			if !wordWrap && len(log) > w-3 {
				log = log[:w-3] + "..."
			}
			lines = append(lines, log)
		}
		logs = append(logs, lines...)
	}

	if len(logs) > h {
		logs = logs[len(logs)-h:]
	}

	for i, log := range logs {
		as.MoveTo(x, y+i)
		log = log + strings.Repeat(" ", w-len(log))
		if strings.HasPrefix(log, "[INFO]") {
			as.Cyan(strings.TrimSpace(strings.TrimPrefix(log, "[INFO]")))
			continue
		}
		if strings.HasPrefix(log, "[WARNING]") {
			as.Orange(strings.TrimSpace(strings.TrimPrefix(log, "[WARNING]")))
			continue
		}
		if strings.HasPrefix(log, "[ERROR]") {
			as.Red(strings.TrimSpace(strings.TrimPrefix(log, "[ERROR]")))
			continue
		}
		as.Text(strings.TrimSpace(log))
	}
	return as
}
