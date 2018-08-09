package main

import "github.com/nsf/termbox-go"
import "math/rand"
import "fmt"

var mod = [3]string{"", "'", "2"}
var x = [2]string{"R", "L"}
var y = [2]string{"U", "D"}
var z = [2]string{"F", "B"}

func scramble() []string {
	var s []string
	curr := -1
	for i := 0; i < 25; i++ {
		next := rand.Intn(3)
		if next != curr {
			switch next {
			case 0:
				s = append(s, x[rand.Intn(2)]+mod[rand.Intn(3)])
			case 1:
				s = append(s, y[rand.Intn(2)]+mod[rand.Intn(3)])
			case 2:
				s = append(s, z[rand.Intn(2)]+mod[rand.Intn(3)])
			}
			curr = next
		} else {
			i--
		}
	}
	return s
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
		fmt.Println(scramble())
	}
}
