package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
		assert.Equal(t, "E", fretboard.strings[0].Root.String())
		assert.Equal(t, "B", fretboard.strings[1].Root.String())
		assert.Equal(t, "G", fretboard.strings[2].Root.String())
		assert.Equal(t, "D", fretboard.strings[3].Root.String())
		assert.Equal(t, "A", fretboard.strings[4].Root.String())
		assert.Equal(t, "E", fretboard.strings[5].Root.String())

		assert.Equal(t, uint(1), fretboard.strings[0].Number)
		assert.Equal(t, uint(2), fretboard.strings[1].Number)
		assert.Equal(t, uint(3), fretboard.strings[2].Number)
		assert.Equal(t, uint(4), fretboard.strings[3].Number)
		assert.Equal(t, uint(5), fretboard.strings[4].Number)
		assert.Equal(t, uint(6), fretboard.strings[5].Number)
	})
}
