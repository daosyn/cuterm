package main

import "github.com/nsf/termbox-go"
import "math/rand"
import "fmt"

var front = [3]string{"F", "F'", "F2"}
var back = [3]string{"B", "B'", "B2"}
var right = [3]string{"R", "R'", "R2"}
var left = [3]string{"L", "L'", "L2"}
var up = [3]string{"U", "U'", "U2"}
var down = [3]string{"D", "D'", "D2"}

func scramble() []string {
	var s []string
	curr := -1
	for i := 0; i < 25; i++ {
		next := rand.Intn(6)
		if next != curr {
			switch next {
			case 0:
				s = append(s, front[rand.Intn(3)])
			case 1:
				s = append(s, back[rand.Intn(3)])
			case 2:
				s = append(s, right[rand.Intn(3)])
			case 3:
				s = append(s, left[rand.Intn(3)])
			case 4:
				s = append(s, up[rand.Intn(3)])
			case 5:
				s = append(s, down[rand.Intn(3)])
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
