package fretboard

import (
	"fmt"
)

const (
	ScaleMinor         string = "minor"
	ScaleHarmonicMinor        = "harmonic minor"
	ScaleMajor                = "major"
)

var (
	notes = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
)

type Scale interface {
	Title() string
	Root() Note
	Contains(Note) bool
}

type scale struct {
	root  Note
	notes []Note
}

func (s scale) Root() Note {
	return s.root
}

func (s scale) Contains(note Note) bool {
	for _, n := range s.notes {
		if n == note {
			return true
		}
	}
	return false
}

type minorScale struct {
	scale
}

func newMinorScale(root Note) minorScale {
	return minorScale{
		scale{
			root:  root,
			notes: buildScaleNotes(root, 2, 3, 5, 7, 8, 10),
		},
	}
}

func (m minorScale) Title() string {
	return fmt.Sprintf("%s minor", m.root)
}

type majorScale struct {
	scale
}

func newMajorScale(root Note) majorScale {
	return majorScale{
		scale{
			root:  root,
			notes: buildScaleNotes(root, 2, 4, 5, 7, 9, 11),
		},
	}
}

func (m majorScale) Title() string {
	return fmt.Sprintf("%s major", m.root)
}

type harmonicMinorScale struct {
	scale
}

func newHarmonicMinorScale(root Note) harmonicMinorScale {
	return harmonicMinorScale{
		scale{
			root:  root,
			notes: buildScaleNotes(root, 2, 3, 5, 7, 8, 11),
		},
	}
}

func (m harmonicMinorScale) Title() string {
	return fmt.Sprintf("%s major", m.root)
}

type emptyScale struct {
}

func (e emptyScale) Title() string {
	return ""
}

func (e emptyScale) Root() Note {
	return Note{value: ""}
}

func (e emptyScale) Contains(_ Note) bool {
	return false
}

func NewScale(rootNote string, scaleType string) (Scale, error) {
	root, err := NewNote(rootNote)
	if err != nil {
		return emptyScale{}, err
	}

	switch scaleType {
	case ScaleMinor:
		return newMinorScale(root), nil
	case ScaleMajor:
		return newMajorScale(root), nil
	case ScaleHarmonicMinor:
		return newHarmonicMinorScale(root), nil
	default:
		return emptyScale{}, fmt.Errorf("scale type %s is not supported", scaleType)
	}
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

func (n Note) Add(halfsteps uint) Note {
	if halfsteps == 0 || halfsteps%12 == 0 {
		return n
	}

	var currentNoteIndex uint = 0
	for i, note := range notes {
		if n.value == note {
			currentNoteIndex = uint(i)
			break
		}
	}

	nextNoteIndex := currentNoteIndex + halfsteps
	if int(nextNoteIndex) > 11 {
		nextNoteIndex = nextNoteIndex % 12
	}

	return Note{value: notes[nextNoteIndex]}
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
