package main

import (
	"flag"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"time"
)

var (
	filename string
	verbose  bool
)

func init() {
	flag.StringVar(&filename, "file", "", "file to process")
	flag.StringVar(&filename, "f", "", "file to process (shorthand)")
	flag.BoolVar(&verbose, "v", false, "enable verbose output. THIS SLOWS EXECUTION DOWN TO AN ABSURD DEGREE. NOT RECOMMENDED")
}

func main() {
	flag.Parse()
	if filename == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	//load image
	fmt.Print("Opening file... ")
	imageFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Success")

	//decode image
	fmt.Print("Decoding image... ")
	imageData, imageType, err := image.Decode(imageFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	imageFile.Close()

	fmt.Printf("Identified type %s\n", imageType)
	notes := GetNotes()
	ratings := make(map[string]int)
	for _, note := range notes {
		ratings[note.name] = 0
	}

	//start iteration
	maxX, maxY, minX, minY := imageData.Bounds().Max.X, imageData.Bounds().Max.Y, imageData.Bounds().Min.X, imageData.Bounds().Min.Y
	width, height := maxX-minX, maxY-minY
	fmt.Printf("Image dimensions: %d x %d. Total iterations to run: %d\n\n", width, height, width*height)
	fmt.Println("Running. This may take a while...")
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
			for _, note := range notes {
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

	fmt.Printf("\nDone in %s!\nResults: %v\n", time.Since(start), ratings)

}

//comment out lines that use this later
func vOut(text string) {
	if verbose {
		fmt.Print(text)
	}
}
