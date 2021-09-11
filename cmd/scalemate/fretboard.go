package main

import (
	"errors"
	"fmt"
)

var ErrNoteNotInScale = errors.New("note is not in scale")

type Fretboard struct {
	tuning  string
	strings []guitarString
	scale   Scale
}

func NewFretboard(tuning string, scale Scale) (Fretboard, error) {
	strings, err := buildStringsFromTuning(tuning)
	if err != nil {
		return Fretboard{}, err
	}

	f := Fretboard{tuning: tuning, strings: strings, scale: scale}
	return f, nil
}

// Fret returns the note on the fretboard for the given string number
// and fret number. It returns a zero-value Note and an error if the string
// number is less than 1 or larger than the number of strings that have been
// parsed from the fretboards tuning. If the Note is valid, but not part of the
// fretboards scale, it returns the Note and ErrNoteNotInScale.
func (f Fretboard) Fret(string, fret uint) (Note, error) {
	if string < 1 || int(string) > len(f.strings) {
		return Note{}, fmt.Errorf("string %d is invalid", string)
	}

	note := f.strings[string-1].fret(fret)
	if !f.scale.Contains(note) {
		return note, ErrNoteNotInScale
	}
	return note, nil
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
