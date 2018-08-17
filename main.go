package main

import (
	"github.com/daosyn/cuterm/pkg/scrambler"
	"github.com/nsf/termbox-go"
	"time"
)

const coldef = termbox.ColorDefault

var scramble []string = scrambler.NewScramble()

var startTime, stopTime time.Time
var times []time.Duration

var width, height int

func setCells(x, y int, msg string, fg, bg termbox.Attribute) int {
    length := 0
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
        length++
		x++
	}
	termbox.Flush()
    return length
}

func getFaceColor(char rune) termbox.Attribute {
	switch char {
	case 'U':
        return termbox.ColorWhite
	case 'L':
		return termbox.ColorMagenta
	case 'F':
		return termbox.ColorGreen
	case 'R':
		return termbox.ColorRed
	case 'B':
		return termbox.ColorBlue
	case 'D':
		return termbox.ColorYellow
	}
	return coldef
}

func drawFace(startx, starty int, face string) {
	i := 0
	fg := termbox.ColorBlack
	for x := startx; x < startx+3; x++ {
		for y := starty; y < starty+3; y++ {
			facelet := rune(face[i])
			color := getFaceColor(facelet)
			termbox.SetCell(x, y, '#', fg, color)
			i++
		}
	}
}

func displayLayout(layout string) {
	// U -> R -> F -> D -> L -> B
	drawFace(3, 0, layout[0:9])
	drawFace(6, 3, layout[9:18])
	drawFace(3, 3, layout[18:27])
	drawFace(3, 6, layout[27:36])
	drawFace(0, 3, layout[36:45])
	drawFace(9, 3, layout[45:54])
	termbox.Flush()
}

func startStopwatch() {
	startTime = time.Now()
	stopTime = time.Time{}
	go func() {
		for stopTime.IsZero() {
			setCells(
				width/2, height/3,
				time.Since(startTime).String(),
				coldef, coldef)
		}
	}()
}

func stopStopwatch() {
	stopTime = time.Now()
	solveTime := stopTime.Sub(startTime)
	times = append(times, solveTime)
	startTime = time.Time{}

    scramble = scrambler.NewScramble()
    termbox.Clear(coldef, coldef)
    initialize()
}

func handleKeyEvent() {
	if startTime.IsZero() {
		startStopwatch()
	} else {
		stopStopwatch()
	}
}

func initialize() {
	width, height = termbox.Size()
	x := width/2 - 30
	y := height / 2
	for _, s := range scramble {
		length := setCells(x, y, s, coldef, coldef)
		x += length + 1
	}
    x = 0
    y = height - 1
    for _, solveTime := range times {
        length := setCells(x, y, solveTime.String(), coldef, coldef)
        x += length + 1
    }
	displayLayout("UUUUUUUUURRRRRRRRRFFFFFFFFFDDDDDDDDDLLLLLLLLLBBBBBBBBB")
}

func mainloop() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return
			}
			if ev.Key == termbox.KeySpace {
				handleKeyEvent()
			}
		case termbox.EventResize:
			termbox.Clear(coldef, coldef)
            initialize()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	initialize()
	mainloop()
}
