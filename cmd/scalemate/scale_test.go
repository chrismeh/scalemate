package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewScale(t *testing.T) {
	t.Run("return error when note does not exist", func(t *testing.T) {
		_, err := NewScale("M", NaturalMinor)
		assert.Error(t, err)
	})

	t.Run("return error when scale type is unknown", func(t *testing.T) {
		_, err := NewScale("A", "Foo")
		assert.Error(t, err)
	})

	t.Run("build correct natural minor scale", func(t *testing.T) {
		scale, err := NewScale("A", NaturalMinor)

		assert.NoError(t, err)
		assert.Equal(t, "A B C D E F G", scale.String())
	})

	t.Run("build correct major scale", func(t *testing.T) {
		scale, err := NewScale("C", Major)

		assert.NoError(t, err)
		assert.Equal(t, "C D E F G A B", scale.String())
	})
}

func TestNewNote(t *testing.T) {
	t.Run("return Note struct with correct value", func(t *testing.T) {
		n, err := NewNote("A")

		assert.NoError(t, err)
		assert.Equal(t, "A", n.String())
	})

	t.Run("return error when note does not exist", func(t *testing.T) {
		_, err := NewNote("M")
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
		{Name: "add 0 halfsteps from A", StartNote: "A", Halfsteps: 0, ExpectedNote: "A"},
		{Name: "add 1 halfstep from A", StartNote: "A", Halfsteps: 1, ExpectedNote: "A#"},
		{Name: "add 2 halfsteps from A", StartNote: "A", Halfsteps: 2, ExpectedNote: "B"},
		{Name: "add 12 halfsteps from A", StartNote: "A", Halfsteps: 12, ExpectedNote: "A"},
		{Name: "add 13 halfsteps from A", StartNote: "A", Halfsteps: 13, ExpectedNote: "A#"},
		{Name: "add 1 halfstep from E", StartNote: "E", Halfsteps: 1, ExpectedNote: "F"},
		{Name: "add 5 halfsteps from E", StartNote: "E", Halfsteps: 5, ExpectedNote: "A"},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			start := Note{value: tt.StartNote}
			result := start.Add(tt.Halfsteps)

			assert.Equal(t, tt.ExpectedNote, result.String())
		})
	}
}
