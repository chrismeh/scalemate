package fretboard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChord_Notes(t *testing.T) {
	t.Run("return root note without any intervals", func(t *testing.T) {
		root := Note{value: "C"}
		c := Chord{root: root, intervals: []uint{}}
		notes := c.Notes()

		assert.Len(t, notes, 1)
		assert.Equal(t, root, notes[0])
	})

	t.Run("return correct notes with a single interval", func(t *testing.T) {
		root := Note{value: "C"}
		c := Chord{root: root, intervals: []uint{4}}
		notes := c.Notes()

		assert.Equal(t, []Note{{value: "C"}, {value: "E"}}, notes)
	})

	t.Run("return correct notes with multiple intervals", func(t *testing.T) {
		root := Note{value: "C"}
		c := Chord{root: root, intervals: []uint{4, 7, 11}}
		notes := c.Notes()

		assert.Equal(t, []Note{{value: "C"}, {value: "E"}, {value: "G"}, {value: "B"}}, notes)
	})
}
