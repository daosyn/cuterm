package main

import (
	"github.com/daosyn/cuterm/pkg/scrambler"
	"github.com/nsf/termbox-go"
	"time"
)

func setCells(x, y int, msg string, fg, bg termbox.Attribute) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
	termbox.Flush()
}

func timer() {
	ticker := time.NewTicker(time.Millisecond)
	go func() {
		for t := range ticker.C {
			setCells(4, 4, t.String(), termbox.ColorDefault, termbox.ColorDefault)
		}
	}()
	time.Sleep(5 * time.Second)
	ticker.Stop()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	scramble := scrambler.NewScramble()
	x := 0 // TODO get x and y from size
	y := 0
	for _, s := range scramble {
		// TODO write to center of screen
		// iterate through each array string
		setCells(x, y, s, termbox.ColorDefault, termbox.ColorDefault)
		x += 3
	}
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}
		}
		timer()
	}
}
