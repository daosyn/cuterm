package main

import (
	"github.com/daosyn/cuterm/pkg/scrambler"
	"github.com/nsf/termbox-go"
	"time"
)

var startTime, stopTime time.Time
var times []time.Duration
var width, height int

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
			setCells(width/3, height/3, time.Since(startTime).String(), termbox.ColorDefault, termbox.ColorDefault)
		}
	}()
}

func stopStopwatch() {
	stopTime = time.Now()
	solveTime := stopTime.Sub(startTime)
	times = append(times, solveTime)
	startTime = time.Time{}
}

func initialize() {
	width, height = termbox.Size()

	x := width/2 - 30
	y := height/2 - 3
	scramble := scrambler.NewScramble()
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
				return
			}
			if ev.Key == termbox.KeySpace {
				handleKeyEvent()
			}
		case termbox.EventResize:
			// TODO draw everything again
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			width, height = termbox.Size()
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
