package main

import (
	"github.com/daosyn/cuterm/pkg/scrambler"
	"github.com/nsf/termbox-go"
	"strconv"
	"time"
)

func setCells(x, y int, msg string, fg, bg termbox.Attribute) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
	termbox.Flush()
}

func startTimer() *time.Ticker {
	// start := time.Now()
	timer := 0.000
	ticker := time.NewTicker(time.Millisecond)
	go func() {
		for range ticker.C {
			timer += 0.001
			setCells(4, 4, strconv.FormatFloat(timer, 'f', -1, 32), termbox.ColorDefault, termbox.ColorDefault)
		}
	}()
	return ticker
}

var ticker *time.Ticker

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
	if ticker == nil {
		ticker = startTimer()
	} else {
		ticker.Stop()
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
