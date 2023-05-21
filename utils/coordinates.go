package utils

import (
	"os"

	"golang.org/x/term"
)

func getTermSize() (int, int) {
	// Retrieve the file descriptor for the standard input (0)
	fd := int(os.Stdin.Fd())

	// Get the terminal size
	width, height, err := term.GetSize(fd)
	if err != nil {
		return -1, -1
	}

	return width, height
}

// Abs converts one-based coordinates to zero-based coordinates.
// Negative x/y values will be calculated as offset from the w/h values.
// Returns the resulting absolute x, y, x2 and y2 coordinates.
func Abs(x, y, w, h int) (int, int, int, int) {
	tw, th := getTermSize()
	if x < 0 {
		x = tw + x - w + 2
	}
	x2 := x + w - 1

	if y < 0 {
		y = th + y - h + 2
	}
	y2 := y + h
	return x, y, x2, y2
}

func AbsX(x int) int {
	if x < 0 {
		tw, _ := getTermSize()
		x = tw + x + 2
	}
	return x
}

func AbsY(y int) int {
	if y < 0 {
		_, th := getTermSize()
		y = th + y + 2
	}
	return y
}
