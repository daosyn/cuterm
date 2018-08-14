package main

import (
	"github.com/daosyn/cuterm/pkg/scrambler"
	"github.com/nsf/termbox-go"
	"time"
)

func timer() {
	ticker := time.NewTicker(time.Millisecond)
	go func() {
		for t := range ticker.C {
			// TODO write to center of screen and update
			termbox.SetCell(4, 4, []rune(t.String())[5], termbox.ColorDefault, termbox.ColorDefault)
			termbox.Flush()
		}
	}()
	time.Sleep(10 * time.Millisecond)
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
	for _, c := range scramble {
		// TODO write to center of screen
		// iterate through each array string
		termbox.SetCell(x, y, []rune(c)[0], termbox.ColorDefault, termbox.ColorDefault)
		x++
	}
	termbox.Flush()
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
