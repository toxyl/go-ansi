package utils

import (
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"golang.org/x/term"
)

var (
	TermWindow = &TerminalWindow{
		Rect: &Rect{
			P1: &Point{
				X: 0,
				Y: 0,
			},
			P2: &Point{
				X: 0,
				Y: 0,
			},
		},
		Cursor: &Point{
			X: 0,
			Y: 0,
		},
	}
)

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	p := &Point{
		X: x,
		Y: y,
	}
	return p
}

type Rect struct {
	P1 *Point
	P2 *Point
}

func (r *Rect) order() {
	p1 := r.P1
	p2 := r.P2

	minX := math.Min(float64(p1.X), float64(p2.X))
	maxX := math.Max(float64(p1.X), float64(p2.X))
	minY := math.Min(float64(p1.Y), float64(p2.Y))
	maxY := math.Max(float64(p1.Y), float64(p2.Y))

	r.P1.X = int(minX)
	r.P2.X = int(maxX)

	r.P1.Y = int(minY)
	r.P2.Y = int(maxY)
}

func (r *Rect) Width() int {
	r.order()
	return r.P2.X - r.P1.X
}

func (r *Rect) Height() int {
	r.order()
	return r.P2.Y - r.P1.Y
}

func (r *Rect) SetWidth(w int) {
	r.order()
	r.P2.X = r.P1.X + w
}

func (r *Rect) SetHeight(h int) {
	r.order()
	r.P2.Y = r.P1.Y + h
}

func (r *Rect) SetX(x int) {
	w := r.Width()
	r.P1.X = x
	r.SetWidth(w)
}

func (r *Rect) SetY(y int) {
	h := r.Height()
	r.P1.Y = y
	r.SetHeight(h)
}

type TerminalWindow struct {
	Rect   *Rect
	Cursor *Point
}

func (tw *TerminalWindow) W() int { return tw.Rect.Width() }
func (tw *TerminalWindow) H() int { return tw.Rect.Height() }

func (tw *TerminalWindow) X() int { return tw.Cursor.X }
func (tw *TerminalWindow) Y() int { return tw.Cursor.Y }

func (tw *TerminalWindow) SetCursorX(x int) {
	if x < 0 {
		x = tw.Rect.Width() + x
	}
	tw.Cursor.X = x
}

func (tw *TerminalWindow) SetCursorY(y int) {
	if y < 0 {
		y = tw.Rect.Height() + y
	}
	tw.Cursor.Y = y
}

func getWindowSize() (int, int, error) {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "mode", "con")
		cmd.Stdin = os.Stdin
		out, err := cmd.Output()
		if err != nil {
			return 0, 0, err
		}
		lines := strings.Split(string(out), "\r\n")
		var cols, rows int
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "Columns:") {
				line = strings.TrimPrefix(line, "Columns:")
				line = strings.TrimSpace(line)
				cols, _ = strconv.Atoi(line)
			}
			if strings.HasPrefix(line, "Lines:") {
				line = strings.TrimPrefix(line, "Lines:")
				line = strings.TrimSpace(line)
				rows, _ = strconv.Atoi(line)
			}
		}
		return cols, rows, nil
	}
	fd := int(os.Stdin.Fd())
	return termGetSize(fd)
}

func termGetSize(fd int) (int, int, error) {
	return term.GetSize(fd)
}

func (tw *TerminalWindow) Update() {
	w, h, err := getWindowSize()
	if err != nil {
		return
	}
	tw.Rect.P1 = NewPoint(0, 0)
	tw.Rect.P2 = NewPoint(w, h)
}

func (tw *TerminalWindow) SelectArea(x1, y1, w, h int) (x int, y int, width int, height int) {
	if x1 < 0 {
		x1 = tw.W() + x1 - 2
	}
	if y1 < 0 {
		y1 = tw.H() + y1 - 2
	}

	x2 := x1 + w
	if w < 0 {
		x2 = tw.W() + w
	} else if w == 0 {
		x2 = tw.W()
	}

	y2 := y1 + h
	if h < 0 {
		y2 = tw.H() + h
	} else if h == 0 {
		y2 = tw.H()
	}

	r := &Rect{
		P1: &Point{
			X: x1,
			Y: y1,
		},
		P2: &Point{
			X: x2,
			Y: y2,
		},
	}
	r.order()
	return r.P1.X, r.P1.Y, r.Width(), r.Height()
}

func init() {
	TermWindow.Update()
	go func() {
		for {
			time.Sleep(3 * time.Second)
			TermWindow.Update()
		}
	}()
}
