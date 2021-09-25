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
	notes                       = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	intervalsMinorScale         = []uint{2, 3, 5, 7, 8, 10}
	intervalsMajorScale         = []uint{2, 4, 5, 7, 9, 11}
	intervalsHarmonicMinorScale = []uint{2, 3, 5, 7, 8, 11}
)

type Scale struct {
	root      Note
	scaleType string
	notes     []Note
}

func NewScale(rootNote string, scaleType string) (Scale, error) {
	root, err := NewNote(rootNote)
	if err != nil {
		return Scale{}, err
	}

	scaleTypeMapping := map[string][]uint{
		ScaleMinor:         intervalsMinorScale,
		ScaleMajor:         intervalsMajorScale,
		ScaleHarmonicMinor: intervalsHarmonicMinorScale,
	}

	intervals, ok := scaleTypeMapping[scaleType]
	if !ok {
		return Scale{}, fmt.Errorf("scale type %s is not supported", scaleType)
	}

	return Scale{root: root, scaleType: scaleType, notes: buildScale(root, intervals...)}, nil
}

func (s Scale) Title() string {
	return fmt.Sprintf("%s %s", s.root.String(), s.scaleType)
}

func (s Scale) Root() Note {
	return s.root
}

func (s Scale) Contains(note Note) bool {
	for _, n := range s.notes {
		if note == n {
			return true
		}
	}
	return false
}

func buildScale(root Note, intervals ...uint) []Note {
	notes := make([]Note, len(intervals)+1)
	notes[0] = root
	for i, interval := range intervals {
		notes[i+1] = root.Add(interval)
	}

	return notes
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
