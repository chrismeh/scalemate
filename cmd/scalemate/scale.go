package main

import (
	"fmt"
	"strings"
)

const (
	NaturalMinor  string = "natural minor"
	HarmonicMinor        = "harmonic minor"
	Major                = "major"
)

var notes = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

type Scale struct {
	root      note
	scaleType string
	notes     []note
}

func NewScale(rootNote string, scaleType string) (Scale, error) {
	root, err := newNote(rootNote)
	if err != nil {
		return Scale{}, err
	}

	switch scaleType {
	case NaturalMinor:
		return buildMinorScale(root), nil
	case HarmonicMinor:
		return buildHarmonicMinorScale(root), nil
	case Major:
		return buildMajorScale(root), nil
	default:
		return Scale{}, fmt.Errorf("scale type %s is not supported", scaleType)
	}
}

func (s Scale) String() string {
	var sb strings.Builder
	for _, n := range s.notes {
		sb.WriteString(n.String() + " ")
	}

	return strings.TrimSpace(sb.String())
}

func buildMinorScale(root note) Scale {
	return Scale{
		root:      root,
		scaleType: NaturalMinor,
		notes:     buildScale(root, 2, 3, 5, 7, 8, 10),
	}
}

func buildMajorScale(root note) Scale {
	return Scale{
		root:      root,
		scaleType: Major,
		notes:     buildScale(root, 2, 4, 5, 7, 9, 11),
	}
}

func buildHarmonicMinorScale(root note) Scale {
	return Scale{
		root:      root,
		scaleType: Major,
		notes:     buildScale(root, 2, 3, 5, 7, 8, 11),
	}
}

func buildScale(root note, intervals ...uint) []note {
	notes := make([]note, len(intervals)+1)
	notes[0] = root
	for i, interval := range intervals {
		notes[i+1] = root.add(interval)
	}

	return notes
}

type note struct {
	value string
}

func newNote(value string) (note, error) {
	for _, n := range notes {
		if value == n {
			return note{value: n}, nil
		}
	}

	return note{}, fmt.Errorf("note does not exist: %s", value)
}

func (n note) add(halfsteps uint) note {
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

	return note{value: notes[nextNoteIndex]}
}

func (n note) String() string {
	return n.value
}
