package main

import (
	"flag"
	"fmt"
	"github.com/chrismeh/scalemate/pkg/fretboard"
	"github.com/chrismeh/scalemate/pkg/renderer"
	"os"
	"strings"
)

func main() {
	scaleFlag := flag.String("scale", "A minor", "Scale you want to generate")
	chordFlag := flag.String("chord", "", "Chord you want to highlight (e. g. Amin7)")
	tuningFlag := flag.String("tuning", "E A D G B E", "Guitar/bass tuning, notes separated by a whitespace")
	fretsFlag := flag.Uint("frets", 12, "Number of frets on the neck")
	fileFlag := flag.String("file", "scale.png", "Filename for saving the PNG")
	flag.Parse()

	scale, err := buildScale(*scaleFlag)
	if err != nil {
		exitWithError(err)
	}

	tuning, err := fretboard.NewTuning(*tuningFlag)
	if err != nil {
		exitWithError(err)
	}

	fb, err := fretboard.New(fretboard.Options{Tuning: tuning, Frets: *fretsFlag})
	if err != nil {
		exitWithError(err)
	}
	fb.HighlightScale(scale)

	if *chordFlag != "" {
		chord, err := fretboard.ParseChord(*chordFlag)
		if err != nil {
			exitWithError(err)
		}
		fb.HighlightChord(chord)
	}

	f, err := os.Create(*fileFlag)
	if err != nil {
		exitWithError(err)
	}
	defer f.Close()

	options := renderer.PNGOptions{FretboardOffsetX: 40.0, FretboardOffsetY: 50.0, DrawTitle: true}
	r := renderer.NewPNGRenderer(fb, options)
	err = r.Render(f)
	if err != nil {
		_ = f.Close()
		exitWithError(err)
	}
}

func buildScale(scale string) (fretboard.Scale, error) {
	firstWhiteSpace := strings.Index(scale, " ")
	rootNote := scale[:firstWhiteSpace]
	scaleType := scale[firstWhiteSpace+1:]

	return fretboard.NewScale(rootNote, scaleType)
}

func exitWithError(e error) {
	fmt.Println("unable to generate scale:", e)
	os.Exit(1)
}
