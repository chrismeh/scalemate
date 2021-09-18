package fretboard

import (
	"errors"
	"fmt"
	"strings"
)

var (
	TuningStandard = "E A D G B E"
)

type Fretboard struct {
	Tuning  Tuning
	Strings uint
	Frets   uint
	strings []guitarString
	scale   Scale
}

type Options struct {
	Tuning Tuning
	Frets  uint
}

func New(options Options) (*Fretboard, error) {
	if options.Tuning.IsZero() {
		t, _ := NewTuning(TuningStandard)
		options.Tuning = t
	}
	if options.Frets == 0 {
		options.Frets = 22
	}

	f := Fretboard{
		Tuning:  options.Tuning,
		Strings: options.Tuning.Strings(),
		Frets:   options.Frets,
		strings: buildStringsFromTuning(options.Tuning),
		scale:   Scale{},
	}
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
	return Fret{
		Number:      fret,
		Note:        note,
		Highlighted: f.scale.contains(note),
		Root:        note == f.scale.root,
	}, nil
}

func (f *Fretboard) String() string {
	if title := f.scale.String(); title != " " {
		return f.scale.String()
	}

	return "Empty fretboard"
}

type Fret struct {
	Number      uint
	Note        note
	Highlighted bool
	Root        bool
}

type Tuning struct {
	notes []note
}

func NewTuning(notes string) (Tuning, error) {
	noteSlice := strings.Split(notes, " ")
	if len(noteSlice) == 0 {
		return Tuning{}, errors.New("notes of the tuning must be separated by a space")
	}

	t := Tuning{notes: make([]note, len(noteSlice))}
	for i, n := range noteSlice {
		note, err := newNote(n)
		if err != nil {
			return Tuning{}, err
		}
		t.notes[i] = note
	}

	return t, nil
}

func (t Tuning) Notes() []string {
	notes := make([]string, len(t.notes))
	for i, n := range t.notes {
		notes[i] = n.String()
	}

	return notes
}

func (t Tuning) Strings() uint {
	return uint(len(t.notes))
}

func (t Tuning) IsZero() bool {
	return len(t.notes) == 0
}

type guitarString struct {
	root   note
	number uint
}

func (g guitarString) fret(fret uint) note {
	return g.root.Add(fret)
}

func buildStringsFromTuning(tuning Tuning) []guitarString {
	guitarStrings := make([]guitarString, tuning.Strings())

	for i := 0; i < int(tuning.Strings()); i++ {
		rootNote := tuning.notes[i]

		stringNumber := int(tuning.Strings()) - i
		guitarStrings[stringNumber-1] = guitarString{
			root:   rootNote,
			number: uint(stringNumber),
		}
	}

	return guitarStrings
}
