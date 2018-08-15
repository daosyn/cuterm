package main

import (
	"github.com/daosyn/cuterm/pkg/scrambler"
	"github.com/nsf/termbox-go"
	"time"
)

var startTime, stopTime time.Time
var times []time.Duration

func setCells(x, y int, msg string, fg, bg termbox.Attribute) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
	termbox.Flush()
}

func startStopwatch() {
	startTime = time.Now()
	stopTime = time.Time{}
	go func() {
		for stopTime.IsZero() {
			setCells(4, 4, time.Since(startTime).String(), termbox.ColorDefault, termbox.ColorDefault)
		}
	}()
}

func stopStopwatch() {
	stopTime = time.Now()
	solveTime := stopTime.Sub(startTime)
	times = append(times, solveTime)
	startTime = time.Time{}
}

func intialize() {
	scramble := scrambler.NewScramble()
	x := 0 // TODO get x and y from size
	y := 0
	for _, s := range scramble {
		// TODO write to center of screen
		setCells(x, y, s, termbox.ColorDefault, termbox.ColorDefault)
		x += 3
	}
}

func handleKeyEvent() {
	if startTime.IsZero() {
		startStopwatch()
	} else {
		stopStopwatch()
	}
}

func mainloop() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				// close application
				return
			}
			if ev.Key == termbox.KeySpace {
				// core logic
				handleKeyEvent()
			}
		case termbox.EventResize:
			// adjust size
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
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
	intialize()
	mainloop()
}
