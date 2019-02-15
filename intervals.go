package main

//Interval : basic interval struct
type Interval struct {
	semitones int
	ratio     [2]int
	name      string
}

//IntervalList : list of intervals, with array index being distance in semitones
var IntervalList = [13]Interval{
	Interval{0, [2]int{1, 1}, "unison"},
	Interval{1, [2]int{25, 24}, "minor 2nd"},
	Interval{2, [2]int{9, 8}, "major 2nd"},
	Interval{3, [2]int{6, 5}, "minor 3rd"},
	Interval{4, [2]int{5, 4}, "major 3rd"},
	Interval{5, [2]int{4, 3}, "perfect 4th"},
	Interval{6, [2]int{45 / 32}, "augmented 4th / diminished 5th"},
	Interval{7, [2]int{3, 4}, "perfect 5th"},
	Interval{8, [2]int{8, 5}, "minor 6th"},
	Interval{9, [2]int{5, 3}, "major 6th"},
	Interval{10, [2]int{9, 5}, "minor 7th"},
	Interval{11, [2]int{15, 8}, "major 7th"},
	Interval{12, [2]int{2, 1}, "octave"},
}

//GetIntervalByName : get interval of 2 notes by their names
func GetIntervalByName(name1, name2 string) Interval {

	var note1, note2 Note
	for _, note := range NoteList {
		if note.name == name1 {
			note1 = note
		} else if note.name == name2 {
			note2 = note
		}
	}

	diff := 0
	if note1.num > note2.num {
		diff = note1.num - note2.num
	} else if note2.num > note1.num {
		diff = note2.num - note1.num
	}

	return IntervalList[diff]
}
