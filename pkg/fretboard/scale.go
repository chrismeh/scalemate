package fretboard

import (
	"fmt"
)

const (
	ScaleMinor string = "minor"
	ScaleMajor        = "major"
)

var (
	notes = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
)

type Scale struct {
	Root      Note
	notes     []Note
	scaleType string
}

func NewScale(rootNote string, scaleType string) (Scale, error) {
	root, err := NewNote(rootNote)
	if err != nil {
		return Scale{}, err
	}

	switch scaleType {
	case ScaleMinor:
		return Scale{Root: root, scaleType: scaleType, notes: buildScaleNotes(root, 2, 3, 5, 7, 8, 10)}, nil
	case ScaleMajor:
		return Scale{Root: root, scaleType: scaleType, notes: buildScaleNotes(root, 2, 4, 5, 7, 9, 11)}, nil
	default:
		return Scale{}, fmt.Errorf("scale type %s is not supported", scaleType)
	}
}

func (s Scale) Name() string {
	if s.scaleType == "" {
		return ""
	}

	return fmt.Sprintf("%s %s", s.Root, s.scaleType)
}

func (s Scale) Contains(note Note) bool {
	for _, n := range s.notes {
		if n.Equals(note) {
			return true
		}
	}
	return false
}

func (s Scale) Chords() []Chord {
	chords := make([]Chord, len(s.notes))
	for i, n := range s.notes {
		chords[i] = NewChord(n, s.buildChordIntervals(n)...)
	}
	return chords
}

func (s Scale) buildChordIntervals(note Note) []uint {
	numberOfThirds := 3
	var addedIntervals uint

	intervals := make([]uint, numberOfThirds)
	for i := 0; i < numberOfThirds; i++ {
		var interval uint = 3
		if s.Contains(note.Add(4)) {
			interval = 4
		}

		note = note.Add(interval)
		addedIntervals += interval
		intervals[i] = addedIntervals
	}

	return intervals
}

type Note struct {
	value string
}

func NewNote(value string) (Note, error) {
	for _, n := range notes {
		if value == n {
			return Note{value: n}, nil
		}
	}

	return Note{}, fmt.Errorf("note does not exist: %s", value)
}

func (n Note) Equals(other Note) bool {
	return n.value == other.value
}

func (n Note) Add(semitones uint) Note {
	if semitones == 0 || semitones%12 == 0 {
		return n
	}
	currentNoteIndex := findNoteIndex(n)

	nextNoteIndex := currentNoteIndex + semitones
	if int(nextNoteIndex) > 11 {
		nextNoteIndex = nextNoteIndex % 12
	}

	return Note{value: notes[nextNoteIndex]}
}

func (n Note) IntervalTo(other Note) uint {
	if n.Equals(other) {
		return 1
	}

	currentNoteIndex := findNoteIndex(n)
	otherNoteIndex := findNoteIndex(other)
	semitones := int(otherNoteIndex) - int(currentNoteIndex)
	if semitones < 0 {
		semitones += 12
	}

	switch semitones {
	case 1, 2:
		return 2
	case 3, 4:
		return 3
	case 5:
		return 4
	case 7:
		return 5
	case 8, 9:
		return 6
	case 10, 11:
		return 7
	}

	return 0
}

func (n Note) String() string {
	return n.value
}

func buildScaleNotes(root Note, intervals ...uint) []Note {
	notes := make([]Note, len(intervals)+1)
	notes[0] = root
	for i, interval := range intervals {
		notes[i+1] = root.Add(interval)
	}

	return notes
}

func findNoteIndex(note Note) uint {
	for i, n := range notes {
		if n == note.value {
			return uint(i)
		}
	}

	panic(fmt.Sprintf("could not determine index for note %v", note))
}
