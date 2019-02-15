package main

import (
	"github.com/lucasb-eyer/go-colorful"
)

//Note : basic note struct
type Note struct {
	//semitones from F
	num int
	//note name
	name string
	//Hz
	frequency float32
	//cm
	wavelength float32
	//color
	color colorful.Color
	//Light struct
	light Light
}

//Light : basic light struct
type Light struct {
	//THz
	frequency float32
	//nm
	wavelength float32
}

//NoteList : array of notes
var NoteList = [13]Note{
	Note{0, "F", 349.2, 98.88, colorful.Color{82.0 / 255.0, 0, 0}, Light{384.0, 780.8}},
	Note{1, "F#/Gb", 370.0, 93.33, colorful.Color{116.0 / 255.0, 0, 0}, Light{406.8, 736.9}},
	Note{2, "G", 392.0, 88.09, colorful.Color{179.0 / 255.0, 0, 0}, Light{431.0, 695.6}},
	Note{3, "G#/Ab", 415.3, 83.15, colorful.Color{238.0 / 255.0, 0, 0}, Light{456.6, 656.5}},
	Note{4, "A", 440.0, 78.48, colorful.Color{255.0 / 255.0, 99.0 / 255.0, 0}, Light{483.8, 619.7}},
	Note{5, "A#/Bb", 466.2, 74.07, colorful.Color{255.0 / 255.0, 236.0 / 255.0, 0}, Light{512.5, 584.9}},
	Note{6, "B", 493.9, 69.92, colorful.Color{153.0 / 255.0, 255.0 / 255.0, 0}, Light{543.0, 552.1}},
	Note{7, "C", 523.2, 65.99, colorful.Color{40.0 / 255.0, 255.0 / 255.0, 0}, Light{575.3, 521.1}},
	Note{8, "C#/Db", 554.4, 62.29, colorful.Color{0, 255.0 / 255.0, 232.0 / 255.0}, Light{609.5, 491.8}},
	Note{9, "D", 587.3, 58.79, colorful.Color{0, 124.0 / 255.0, 255.0 / 255.0}, Light{645.8, 464.2}},
	Note{10, "D#/Eb", 622.2, 55.49, colorful.Color{5.0 / 255.0, 0, 255.0 / 255.0}, Light{684.2, 438.2}},
	Note{11, "E", 659.3, 52.38, colorful.Color{69.0 / 255.0, 0, 234.0 / 255.0}, Light{724.9, 413.6}},
	Note{12, "F5", 698.5, 49.44, colorful.Color{87.0 / 255.0, 0, 158.0 / 255.0}, Light{768.0, 390.4}},
}
