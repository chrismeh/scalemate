package fretboard

import (
	"fmt"
	"reflect"
)

var (
	intervalsMajor7          = []uint{4, 7, 11}
	intervalsMinor7          = []uint{3, 7, 10}
	intervalsDominant7       = []uint{4, 7, 10}
	intervalsHalfDiminished7 = []uint{3, 6, 10}
)

func NewChord(rootNote Note, intervals ...uint) Chord {
	return Chord{
		Name:  fmt.Sprintf("%s%s", rootNote, identifyChord(intervals)),
		notes: buildChordNotes(rootNote, intervals...),
	}
}

type Chord struct {
	Name  string
	notes []Note
}

func (c Chord) Contains(n Note) bool {
	for _, note := range c.notes {
		if note == n {
			return true
		}
	}
	return false
}

func buildChordNotes(root Note, intervals ...uint) []Note {
	notes := make([]Note, len(intervals)+1)
	notes[0] = root

	for i, v := range intervals {
		notes[i+1] = root.Add(v)
	}

	return notes
}

func identifyChord(intervals []uint) string {
	var suffix string
	switch {
	case reflect.DeepEqual(intervals, intervalsMajor7):
		suffix = "maj7"
	case reflect.DeepEqual(intervals, intervalsMinor7):
		suffix = "min7"
	case reflect.DeepEqual(intervals, intervalsDominant7):
		suffix = "7"
	case reflect.DeepEqual(intervals, intervalsHalfDiminished7):
		suffix = "min7b5"
	default:
		suffix = ""
	}

	return suffix
}
