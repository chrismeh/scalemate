package fretboard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScale_Contains(t *testing.T) {
	t.Run("return true if scale contains note", func(t *testing.T) {
		scale, _ := NewScale("A", ScaleNaturalMinor)
		assert.True(t, scale.contains(note{value: "C"}))
	})

	t.Run("return false if scale does not contain note", func(t *testing.T) {
		scale, _ := NewScale("A", ScaleNaturalMinor)
		assert.False(t, scale.contains(note{value: "C#"}))
	})
}

func TestNewScale(t *testing.T) {
	t.Run("return error when note does not exist", func(t *testing.T) {
		_, err := NewScale("M", ScaleNaturalMinor)
		assert.Error(t, err)
	})

	t.Run("return error when scale type is unknown", func(t *testing.T) {
		_, err := NewScale("A", "Foo")
		assert.Error(t, err)
	})

	t.Run("build correct natural minor scale", func(t *testing.T) {
		scale, err := NewScale("A", ScaleNaturalMinor)

		assert.NoError(t, err)
		assert.Equal(t, "A B C D E F G", scale.String())
	})

	t.Run("build correct major scale", func(t *testing.T) {
		scale, err := NewScale("C", ScaleMajor)

		assert.NoError(t, err)
		assert.Equal(t, "C D E F G A B", scale.String())
	})

	t.Run("build correct harmonic minor scale", func(t *testing.T) {
		scale, err := NewScale("A", ScaleHarmonicMinor)

		assert.NoError(t, err)
		assert.Equal(t, "A B C D E F G#", scale.String())
	})
}

func TestNewNote(t *testing.T) {
	t.Run("return note struct with correct value", func(t *testing.T) {
		n, err := newNote("A")

		assert.NoError(t, err)
		assert.Equal(t, "A", n.String())
	})

	t.Run("return error when note does not exist", func(t *testing.T) {
		_, err := newNote("M")
		assert.Error(t, err)
	})
}

func TestNote_Add(t *testing.T) {
	tests := []struct {
		Name         string
		StartNote    string
		Halfsteps    uint
		ExpectedNote string
	}{
		{Name: "Add 0 halfsteps from A", StartNote: "A", Halfsteps: 0, ExpectedNote: "A"},
		{Name: "Add 1 halfstep from A", StartNote: "A", Halfsteps: 1, ExpectedNote: "A#"},
		{Name: "Add 2 halfsteps from A", StartNote: "A", Halfsteps: 2, ExpectedNote: "B"},
		{Name: "Add 12 halfsteps from A", StartNote: "A", Halfsteps: 12, ExpectedNote: "A"},
		{Name: "Add 13 halfsteps from A", StartNote: "A", Halfsteps: 13, ExpectedNote: "A#"},
		{Name: "Add 1 halfstep from E", StartNote: "E", Halfsteps: 1, ExpectedNote: "F"},
		{Name: "Add 5 halfsteps from E", StartNote: "E", Halfsteps: 5, ExpectedNote: "A"},
		{Name: "Add 22 halfsteps from E", StartNote: "E", Halfsteps: 22, ExpectedNote: "D"},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			start := note{value: tt.StartNote}
			result := start.Add(tt.Halfsteps)

			assert.Equal(t, tt.ExpectedNote, result.String())
		})
	}
}