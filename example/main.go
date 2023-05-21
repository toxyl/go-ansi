package main

import (
	"os"
	"os/signal"
	"syscall"

	ga "github.com/toxyl/go-ansi"
)

func main() {
	// Create a channel to receive the signals
	signalChannel := make(chan os.Signal, 1)

	// Notify the signal channel when an interrupt signal is received (SIGINT)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT)

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
		Box(2, 2, 40, 3, "Hello").
		Box(3, 3, 40, 3, "Hello").
		Box(4, 4, 40, 3, "Hello").
		DarkGreen().
		Box(50, 3, 40, 3, "").
		BoxCustom(
			30, 4, 40, 2,
			"Test",
			"█", "█", "█",
			"█", "░", "█",
			"▀", "▀", "▀",
		).
		MoveHome().
		Fill(ga.DarkRed, 1, 1, 3, 3).
		Fill(ga.DarkGreen, 2, 2, 3, 3).
		Fill(ga.DarkYellow, 3, 3, 3, 3).
		Fill(ga.Red, -1, -1, 3, 3).
		Fill(ga.Green, -2, -2, 3, 3).
		Fill(ga.Yellow, -3, -3, 3, 3).
		Print()

	// Block the main goroutine until a signal is received
	<-signalChannel

	ga.New().MoveHome().Clear().ClearScreen().Println()
}
