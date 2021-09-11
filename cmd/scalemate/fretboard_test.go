package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFretboard_Fret(t *testing.T) {
	scale, _ := NewScale("A", Major)
	fretboard, _ := NewFretboard("EADGBE", scale)

	t.Run("return error for an invalid string number", func(t *testing.T) {
		tests := []struct {
			Name   string
			String uint
		}{
			{Name: "string is zero", String: 0},
			{Name: "string does not exist", String: 7},
		}

		for _, tt := range tests {
			t.Run(tt.Name, func(t *testing.T) {
				_, err := fretboard.Fret(tt.String, 0)
				assert.Error(t, err)
			})
		}
	})

	t.Run("return correct note for specified fret and tuning", func(t *testing.T) {
		note, err := fretboard.Fret(6, 5)

		assert.NoError(t, err)
		assert.Equal(t, "A", note.String())
	})

	t.Run("return correct note and error if note is not in the fretboards scale", func(t *testing.T) {
		note, err := fretboard.Fret(6, 1)

		assert.Error(t, err)
		assert.Equal(t, "F", note.String())
	})
}

func TestNewFretboard(t *testing.T) {
	t.Run("return error when tuning contains a non existing note", func(t *testing.T) {
		scale, _ := NewScale("A", Major)
		_, err := NewFretboard("ZADGBE", scale)

		assert.Error(t, err)
	})

	t.Run("parse guitar strings from specified tuning", func(t *testing.T) {
		scale, _ := NewScale("A", Major)
		fretboard, err := NewFretboard("EADGBE", scale)

		assert.NoError(t, err)
		assert.Equal(t, "E", fretboard.strings[0].root.String())
		assert.Equal(t, "B", fretboard.strings[1].root.String())
		assert.Equal(t, "G", fretboard.strings[2].root.String())
		assert.Equal(t, "D", fretboard.strings[3].root.String())
		assert.Equal(t, "A", fretboard.strings[4].root.String())
		assert.Equal(t, "E", fretboard.strings[5].root.String())

		assert.Equal(t, uint(1), fretboard.strings[0].number)
		assert.Equal(t, uint(2), fretboard.strings[1].number)
		assert.Equal(t, uint(3), fretboard.strings[2].number)
		assert.Equal(t, uint(4), fretboard.strings[3].number)
		assert.Equal(t, uint(5), fretboard.strings[4].number)
		assert.Equal(t, uint(6), fretboard.strings[5].number)
	})
}
