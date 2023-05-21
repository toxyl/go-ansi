package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	ga "github.com/toxyl/go-ansi"
)

func main() {
	// Create a channel to receive the signals
	signalChannel := make(chan os.Signal, 1)

	// Notify the signal channel when an interrupt signal is received (SIGINT)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGINT)
	go func() {
		screen := ga.New()
		log := []string{
			"Elvis has entered the building.",
			"The building has collapsed.",
			"Elvis is dead, time to panic!",
		}
		for {
			screen.
				CursorHide().
				MoveHome().ClearScreen().
				// ClearCursorToScreenEnd().
				// Bold("MY TEST APP ").Green("v0.0.1").
				// Blue().Text(" and ").Italic("italic").
				// Gray().Red(" red").Text(" and ").
				// FGClear().
				// Underline("underlined").
				// BGClear().
				// Yellow().
				// Text(" and ").
				// StrikeThrough("crossed out").
				// MoveToLastRow().
				// Red("This is the last row!").
				// Box(2, 2, 40, 3, "Hello").
				// Box(3, 3, 40, 3, "Hello").
				// Box(4, 4, 40, 3, "Hello").
				// DarkGreen().
				// Box(7, 6, -6, 0, "OUTPUT").
				Box(1, 1, -1, 3, "INPUT").
				LogBox(1, 5, -1, -1, "Log", log, false).
				// Yellow().
				// BoxCustom(
				// 	30, 8, 40, 5,
				// 	"Test",
				// 	"█", "█", "█",
				// 	"█", "░", "█",
				// 	"▀", "▀", "▀",
				// ).
				// Fill(ga.DarkRed, 1, 1, 3, 3).
				// Fill(ga.DarkGreen, 2, 2, 3, 3).
				// Fill(ga.DarkYellow, 3, 3, 3, 3).
				// Fill(ga.Red, -1, -1, 3, 3).
				// Fill(ga.Green, -2, -2, 3, 3).
				// Fill(ga.Yellow, -3, -3, 3, 3).
				MoveTo(2, 2).
				CursorShow().
				Print()

			time.Sleep(1 * time.Second)
			t := time.Now().Unix()
			rnd := ""
			if t%2 == 0 {
				rnd = "[INFO] Hello world?"
			} else if t%3 == 0 {
				rnd = "[WARNING] Something ain't right."
			} else if t%7 == 0 {
				rnd = "[ERROR] You hit the 7th!"
			}
			log = append(log, fmt.Sprintf("%s It is now %s", rnd, time.Now().Format(time.RFC3339Nano)))
		}
	}()

	// Block the main goroutine until a signal is received
	<-signalChannel
	ga.New().MoveHome().Clear().ClearScreen().Println()
}
