package main

import (
	"github.com/nsf/termbox-go"
)

var tDef = termbox.ColorDefault

//Pos : text position type
type Pos int

//
const (
	PosTitle Pos = iota
	PosBody
)

//InitDisplay : start display
func InitDisplay() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	drawBase()

}

func drawBase() {
	if err := termbox.Clear(tDef, tDef); err != nil {
		panic(err)
	}
	width, height := termbox.Size()
	for _, y := range [2]int{0, height - 1} {
		for x := 0; x < width; x++ {
			termbox.SetCell(x, y, '█', tDef, tDef)
		}
	}
	for _, x := range [2]int{0, width - 1} {
		for y := 0; y < height; y++ {
			termbox.SetCell(x, y, '█', tDef, tDef)
		}
	}
	termbox.Flush()

}

//WriteLine : output line
func WriteLine(str string, x int, y int) {
	line := []rune(str)
	for i := 0; i < len(line); i++ {
		termbox.SetCell(x+i, y, line[i], tDef, tDef)
	}
}

//PastLines : output line history
func PastLines(x, y int) { //draw history and flush
	drawBase()
	for i, l := range lineBuffer {
		WriteLine(l, x, y-1-i)
	}
}

//AddToBuffer : add to line history
func AddToBuffer(line string) {
	lineBuffer = append([]string{line}, lineBuffer...)
	if len(lineBuffer) > dy-2 { //clear lines that scroll off the window from memory
		lineBuffer = lineBuffer[:dy]
	}
}
