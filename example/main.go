package main

import (
	"fmt"

	ga "github.com/toxyl/go-ansi"
)

func main() {
	ga.New().
		MoveToFirstRow().ClearCursorToScreenEnd().
		Bold("This is bold").
		Blue().Text(" and ").Italic("italic").
		Gray().Red(" red").Text(" and ").
		FGClear().
		Underline("underlined").
		BGClear().
		Yellow().
		Text(" and ").
		StrikeThrough("crossed out").
		MoveToLastRow().
		Red("This is the last row!").
		Box(0, 1, 40, 3, "Hello").
		DarkGreen().
		Box(50, 3, 40, 3, "").
		Println()
	fmt.Println("")
	fmt.Println("")
}
