package main

import "github.com/nsf/termbox-go"
import "math/rand"
import "fmt"

var move = [18]string{
    "F", "F'", "F2",
    "R", "R'", "R2",
    "L", "L'", "L2",
    "U", "U'", "U2",
    "D", "D'", "D2",
    "B", "B'", "B2",
}

func scramble() []string {
    var s []string
    for i := 0; i < 30; i ++ {
        s = append(s, move[rand.Intn(18)])
        // add logic for redundancies
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
mainloop:
    for {
        switch ev := termbox.PollEvent(); ev.Type {
        case termbox.EventKey:
            if ev.Key == termbox.KeyEsc {
                break mainloop
            }
        }
        fmt.Println(scramble())
    }
}
