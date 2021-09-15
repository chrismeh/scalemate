package main

import (
	"fmt"
)

type Fretboard struct {
	tuning  string
	strings []guitarString
	scale   Scale
}

func NewFretboard(tuning string) (*Fretboard, error) {
	strings, err := buildStringsFromTuning(tuning)
	if err != nil {
		return nil, err
	}

	f := Fretboard{tuning: tuning, strings: strings, scale: Scale{}}
	return &f, nil
}

func (f *Fretboard) HighlightScale(s Scale) {
	f.scale = s
}

func (f *Fretboard) Fret(string, fret uint) (Fret, error) {
	if string < 1 || int(string) > len(f.strings) {
		return Fret{}, fmt.Errorf("string %d is invalid", string)
	}

	note := f.strings[string-1].fret(fret)
	return Fret{Number: fret, Note: note, Highlighted: f.scale.Contains(note)}, nil
}

type Fret struct {
	Number      uint
	Note        Note
	Highlighted bool
}

type guitarString struct {
	root   Note
	number uint
}

func (g guitarString) fret(fret uint) Note {
	return g.root.Add(fret)
}

func buildStringsFromTuning(tuning string) ([]guitarString, error) {
	strings := make([]guitarString, len(tuning))
	for i := 0; i < len(tuning); i++ {
		rootNote, err := NewNote(string(tuning[i]))
		if err != nil {
			return nil, err
		}

		stringNumber := len(tuning) - i
		strings[stringNumber-1] = guitarString{
			root:   rootNote,
			number: uint(stringNumber),
		}
	}

	return strings, nil
}
