package goansi

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
