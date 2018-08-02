package main

import "github.com/nsf/termbox-go"
import "math/rand"
import "fmt"

var moves = [18]string{
    "F", "F'", "F2",
    "R", "R'", "R2",
    "L", "L'", "L2",
    "U", "U'", "U2",
    "D", "D'", "D2",
    "B", "B'", "B2",
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
        for i := 0; i < 25; i++ {
            move := rand.Intn(18)
            fmt.Print(moves[move]+" ")
        }
        fmt.Println()
    }
}
