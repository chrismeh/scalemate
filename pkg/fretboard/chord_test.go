package fretboard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewChord(t *testing.T) {
	root := Note{value: "C"}

	t.Run("set correct name for each chord", func(t *testing.T) {
		tests := []struct {
			Name         string
			Intervals    []uint
			ExpectedName string
		}{
			{Name: "major 7", Intervals: intervalsMajor7, ExpectedName: "Cmaj7"},
			{Name: "minor 7", Intervals: intervalsMinor7, ExpectedName: "Cmin7"},
			{Name: "dominant 7", Intervals: intervalsDominant7, ExpectedName: "C7"},
			{Name: "half diminished 7", Intervals: intervalsHalfDiminished7, ExpectedName: "Cmin7b5"},
		}

		for _, tt := range tests {
			c := NewChord(root, tt.Intervals...)
			assert.Equal(t, tt.ExpectedName, c.Name)
		}
	})

	t.Run("set correct notes for each chord", func(t *testing.T) {
		tests := []struct {
			Name          string
			Intervals     []uint
			ExpectedNotes []string
		}{
			{Name: "major 7", Intervals: intervalsMajor7, ExpectedNotes: []string{"C", "E", "G", "B"}},
			{Name: "minor 7", Intervals: intervalsMinor7, ExpectedNotes: []string{"C", "D#", "G", "A#"}},
			{Name: "dominant 7", Intervals: intervalsDominant7, ExpectedNotes: []string{"C", "E", "G", "A#"}},
			{Name: "half diminished 7", Intervals: intervalsHalfDiminished7, ExpectedNotes: []string{"C", "D#", "F#", "A#"}},
		}

		for _, tt := range tests {
			t.Run(tt.Name, func(t *testing.T) {
				c := NewChord(root, tt.Intervals...)
				for _, n := range tt.ExpectedNotes {
					assert.True(t, c.Contains(Note{value: n}))
				}
			})
		}
	})
}
