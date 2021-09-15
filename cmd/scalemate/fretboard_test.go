package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFretboard_Fret(t *testing.T) {
	scale, _ := NewScale("A", Major)

	t.Run("return error for an invalid string number", func(t *testing.T) {
		fretboard, _ := NewFretboard(FretboardOptions{Tuning: "EADGBE"})

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

	t.Run("return correct fret for specified fret number and tuning", func(t *testing.T) {
		fretboard, _ := NewFretboard(FretboardOptions{Tuning: "EADGBE"})
		fretboard.HighlightScale(scale)
		fret, err := fretboard.Fret(6, 5)

		assert.NoError(t, err)
		assert.Equal(t, uint(5), fret.Number)
		assert.Equal(t, "A", fret.Note.String())
		assert.Equal(t, true, fret.Highlighted)
	})

	t.Run("return false for Highlighted if a frets note is not in scale", func(t *testing.T) {
		fretboard, _ := NewFretboard(FretboardOptions{Tuning: "EADGBE"})
		fretboard.HighlightScale(scale)
		fret, err := fretboard.Fret(6, 1)

		assert.NoError(t, err)
		assert.Equal(t, false, fret.Highlighted)
	})

	t.Run("return false for Highlighted if the fretboard has no highlighted scale", func(t *testing.T) {
		fretboard, _ := NewFretboard(FretboardOptions{Tuning: "EADGBE"})
		fret, err := fretboard.Fret(6, 1)

		assert.NoError(t, err)
		assert.Equal(t, false, fret.Highlighted)
	})
}

func TestNewFretboard(t *testing.T) {
	t.Run("use EADGBE tuning for zero value FretboardOptions", func(t *testing.T) {
		fretboard, err := NewFretboard(FretboardOptions{})

		assert.NoError(t, err)
		assert.Equal(t, uint(6), fretboard.Strings)
		assert.Equal(t, "E", fretboard.strings[0].root.String())
		assert.Equal(t, "B", fretboard.strings[1].root.String())
		assert.Equal(t, "G", fretboard.strings[2].root.String())
		assert.Equal(t, "D", fretboard.strings[3].root.String())
		assert.Equal(t, "A", fretboard.strings[4].root.String())
		assert.Equal(t, "E", fretboard.strings[5].root.String())
	})

	t.Run("use 22 frets as default for zero value FretboardOptions", func(t *testing.T) {
		fretboard, err := NewFretboard(FretboardOptions{})

		assert.NoError(t, err)
		assert.Equal(t, uint(22), fretboard.Frets)
	})

	t.Run("return error when tuning contains a non existing note", func(t *testing.T) {
		_, err := NewFretboard(FretboardOptions{Tuning: "ZADGBE"})

		assert.Error(t, err)
	})

	t.Run("parse guitar strings from specified tuning", func(t *testing.T) {
		fretboard, err := NewFretboard(FretboardOptions{Tuning: "DADGBE"})

		assert.NoError(t, err)
		assert.Equal(t, "E", fretboard.strings[0].root.String())
		assert.Equal(t, "B", fretboard.strings[1].root.String())
		assert.Equal(t, "G", fretboard.strings[2].root.String())
		assert.Equal(t, "D", fretboard.strings[3].root.String())
		assert.Equal(t, "A", fretboard.strings[4].root.String())
		assert.Equal(t, "D", fretboard.strings[5].root.String())

		assert.Equal(t, uint(1), fretboard.strings[0].number)
		assert.Equal(t, uint(2), fretboard.strings[1].number)
		assert.Equal(t, uint(3), fretboard.strings[2].number)
		assert.Equal(t, uint(4), fretboard.strings[3].number)
		assert.Equal(t, uint(5), fretboard.strings[4].number)
		assert.Equal(t, uint(6), fretboard.strings[5].number)
	})
}
