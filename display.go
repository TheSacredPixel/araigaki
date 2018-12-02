package main

import (
	"github.com/nsf/termbox-go"
)

var tDef = termbox.ColorDefault

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
			termbox.SetCell(x, y, '—', tDef, tDef)
		}
	}
	for _, x := range [2]int{0, width - 1} {
		for y := 0; y < height; y++ {
			termbox.SetCell(x, y, '|', tDef, tDef)
		}
	}
	termbox.Flush()

}
