package main

import (
	"flag"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/nsf/termbox-go"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"sort"
	"time"
)

var (
	filename      string
	verbose, lite bool
	dx, dy        int
	lineBuffer    []string
)

type result struct {
	name string
	num  int
}

func init() {
	flag.StringVar(&filename, "file", "", "file to process")
	flag.StringVar(&filename, "f", "", "file to process (shorthand)")
	flag.BoolVar(&lite, "l", true, "lite output") //TEMP true
	flag.BoolVar(&verbose, "v", false, "enable verbose output. THIS SLOWS EXECUTION DOWN TO AN ABSURD DEGREE. FOR TESTING ONLY")
}

func main() {
	flag.Parse()
	if filename == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}
	if !lite {
		InitDisplay()
		defer termbox.Close()
		dx, dy = termbox.Size()
	}

	//load image
	out("Opening file... ", PosBody)
	imageFile, err := os.Open(filename)
	if err != nil {
		out(fmt.Sprintf("%s\n", err), PosBody)
		os.Exit(1)
	}
	out("Success\n", PosBody)

	//decode image
	out("Decoding image... ", PosBody)
	imageData, imageType, err := image.Decode(imageFile)
	if err != nil {
		out(fmt.Sprintf("%s\n", err), PosBody)
		os.Exit(1)
	}
	imageFile.Close()

	out(fmt.Sprintf("Identified type %s\n", imageType), PosBody)
	ratings := make(map[string]int)
	for _, note := range NoteList {
		ratings[note.name] = 0
	}

	//start iteration
	maxX, maxY, minX, minY := imageData.Bounds().Max.X, imageData.Bounds().Max.Y, imageData.Bounds().Min.X, imageData.Bounds().Min.Y
	width, height := maxX-minX, maxY-minY
	out(fmt.Sprintf("Image dimensions: %d x %d. Total iterations to run: %d\n\n", width, height, width*height), PosBody)
	out("Running. This may take a while...\n", PosBody)
	start := time.Now()
	for y := 0; y < imageData.Bounds().Max.Y; y++ {
		for x := 0; x < imageData.Bounds().Max.X; x++ {
			//get color at pixel x,y
			pixel := imageData.At(minX+x, minY+y)
			r, g, b, _ := pixel.RGBA()
			vOut(fmt.Sprintf("At %d, %d: %d, %d, %d ", minX+x, minY+y, r, g, b))
			//convert to colorful.Color
			color, ok := colorful.MakeColor(pixel)
			if !ok { //alpha channel is 0
				vOut("skipped. ")
				continue
			}
			vOut("converted. ")

			//compare to notes and find closest
			minDistance := 5.0
			var closestNote Note
			for _, note := range NoteList {
				if dist := color.DistanceLab(note.color); dist < minDistance {
					vOut(fmt.Sprintf("Distance from %s: %f ", note.name, dist))
					closestNote = note
					minDistance = dist
				}
			}

			//increment rating
			ratings[closestNote.name]++
			vOut(fmt.Sprintf("Closest is %s\n", closestNote.name))
		}
	}

	//parse results
	var results []result
	for key, value := range ratings {
		results = append(results, result{key, value})
	}

	//sort results descending
	sort.Slice(results, func(i, j int) bool {
		return results[i].num > results[j].num
	})

	out(fmt.Sprintf("\nDone in %s!\nResults: %v\n", time.Since(start), results), PosBody)

	//primary interval
	intr := GetIntervalByName(results[0].name, results[1].name)

	out(fmt.Sprintf("Interval: %v\n", intr), PosBody)

}

func out(text string, pos Pos) {
	if lite {
		fmt.Print(text)
	} else {
		x, y := 0, 0
		switch pos {
		case PosBody:
			x = 1
			y = dy - 2
		}
		PastLines(x, y)
		WriteLine(text, x, y)
		termbox.Flush()
		AddToBuffer(text)
	}
}

//TODO comment out lines that use this later
func vOut(text string) {
	if verbose {
		fmt.Print(text)
	}
}
