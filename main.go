package main

import (
	"fmt"
	"github.com/daosyn/cuterm/pkg/scrambler"
	"github.com/nsf/termbox-go"
	"time"
)

func timer() {
	ticker := time.NewTicker(time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Current time: ", t)
		}
	}()
	time.Sleep(50 * time.Second)
	ticker.Stop()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}
		}
		fmt.Println(scrambler.NewScramble())
		timer()
	}
}
