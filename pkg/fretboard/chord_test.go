package fretboard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChord_Contains(t *testing.T) {
	c := &Chord{root: Note{value: "C"}, intervals: intervalsMajor7}

	assert.True(t, c.Contains(Note{value: "C"}))
	assert.True(t, c.Contains(Note{value: "E"}))
	assert.True(t, c.Contains(Note{value: "G"}))
	assert.True(t, c.Contains(Note{value: "B"}))
}

func TestChord_Name(t *testing.T) {
	root := Note{value: "B"}

	tests := []struct {
		Name              string
		Intervals         []uint
		ExpectedChordName string
	}{
		{Name: "major 7", Intervals: intervalsMajor7, ExpectedChordName: "Bmaj7"},
		{Name: "minor 7", Intervals: intervalsMinor7, ExpectedChordName: "Bmin7"},
		{Name: "dominant 7", Intervals: intervalsDominant7, ExpectedChordName: "B7"},
		{Name: "half-diminished 7", Intervals: intervalsHalfDiminished7, ExpectedChordName: "Bmin7b5"},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			c := Chord{root: root, intervals: tt.Intervals}
			assert.Equal(t, tt.ExpectedChordName, c.Name())
		})
	}
}

func TestChord_buildNotes(t *testing.T) {
	t.Run("return root note without any intervals", func(t *testing.T) {
		root := Note{value: "C"}
		c := Chord{root: root, intervals: []uint{}}
		notes := c.buildNotes()

		assert.Len(t, notes, 1)
		assert.Equal(t, root, notes[0])
	})

	t.Run("return correct notes with a single interval", func(t *testing.T) {
		root := Note{value: "C"}
		c := Chord{root: root, intervals: []uint{4}}
		notes := c.buildNotes()

		assert.Equal(t, []Note{{value: "C"}, {value: "E"}}, notes)
	})

	t.Run("return correct notes with multiple intervals", func(t *testing.T) {
		root := Note{value: "C"}
		c := Chord{root: root, intervals: []uint{4, 7, 11}}
		notes := c.buildNotes()

		assert.Equal(t, []Note{{value: "C"}, {value: "E"}, {value: "G"}, {value: "B"}}, notes)
	})
}
