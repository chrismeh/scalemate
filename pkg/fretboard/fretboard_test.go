package fretboard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTuning(t *testing.T) {
	t.Run("return error when tuning contains an invalid note", func(t *testing.T) {
		_, err := NewTuning("Z A D G B E")
		assert.Error(t, err)
	})

	t.Run("return no error when tuning contains only valid notes", func(t *testing.T) {
		_, err := NewTuning("E A D G B E")
		assert.NoError(t, err)
	})
}

func TestTuning_Strings(t *testing.T) {
	tuning, _ := NewTuning("E A D G B E")

	assert.Equal(t, uint(6), tuning.Strings())
}

func TestTuning_Notes(t *testing.T) {
	tuning, _ := NewTuning("E A D G B E")
	expectedNotes := []Note{
		{value: "E"},
		{value: "A"},
		{value: "D"},
		{value: "G"},
		{value: "B"},
		{value: "E"},
	}

	assert.Equal(t, expectedNotes, tuning.Notes())
}

func TestTuning_IsZero(t *testing.T) {
	t.Run("return true for a zero value struct", func(t *testing.T) {
		tuning := Tuning{}
		assert.Equal(t, true, tuning.IsZero())
	})

	t.Run("return false for an initialised struct", func(t *testing.T) {
		tuning, _ := NewTuning(TuningStandard)
		assert.Equal(t, false, tuning.IsZero())
	})
}

func TestFretboard_String(t *testing.T) {
	t.Run("return the scales title if a scale has been set", func(t *testing.T) {
		scale, _ := NewScale("A", ScaleMajor)
		fb, _ := New(Options{})
		fb.HighlightScale(scale)

		assert.Equal(t, scale.Name(), fb.String())
	})

	t.Run("return a default title if no scale has been set", func(t *testing.T) {
		fb, _ := New(Options{})

		assert.Equal(t, "Empty fretboard", fb.String())
	})
}

func TestFretboard_Fret(t *testing.T) {
	testScale, _ := NewScale("A", ScaleMajor)

	t.Run("return error for an invalid string number", func(t *testing.T) {
		tuning, _ := NewTuning(TuningStandard)
		fretboard, _ := New(Options{Tuning: tuning})

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
		tuning, _ := NewTuning(TuningStandard)
		fretboard, _ := New(Options{Tuning: tuning})
		fretboard.HighlightScale(testScale)
		fret, err := fretboard.Fret(6, 5)

		assert.NoError(t, err)
		assert.Equal(t, uint(5), fret.Number)
		assert.Equal(t, Note{value: "A"}, fret.Note)
		assert.Equal(t, true, fret.Highlighted)
	})

	t.Run("return false for Highlighted if a frets note is not in scale", func(t *testing.T) {
		tuning, _ := NewTuning(TuningStandard)
		fretboard, _ := New(Options{Tuning: tuning})
		fretboard.HighlightScale(testScale)
		fret, err := fretboard.Fret(6, 1)

		assert.NoError(t, err)
		assert.Equal(t, false, fret.Highlighted)
	})

	t.Run("return false for Highlighted if the fretboard has no highlighted scale", func(t *testing.T) {
		tuning, _ := NewTuning(TuningStandard)
		fretboard, _ := New(Options{Tuning: tuning})
		fret, err := fretboard.Fret(6, 1)

		assert.NoError(t, err)
		assert.Equal(t, false, fret.Highlighted)
	})

	t.Run("return true is the frets note is the root note of the highlighted scale", func(t *testing.T) {
		tuning, _ := NewTuning(TuningStandard)
		fretboard, _ := New(Options{Tuning: tuning})
		scale, err := NewScale("A", ScaleMinor)
		fretboard.HighlightScale(scale)

		fret, err := fretboard.Fret(6, 5)

		assert.NoError(t, err)
		assert.Equal(t, true, fret.Root)
	})
}

func TestNewFretboard(t *testing.T) {
	t.Run("use EADGBE tuning for zero value Options", func(t *testing.T) {
		fretboard, err := New(Options{})

		assert.NoError(t, err)
		assert.Equal(t, uint(6), fretboard.Strings)
		assert.Equal(t, "E", fretboard.strings[0].root.String())
		assert.Equal(t, "B", fretboard.strings[1].root.String())
		assert.Equal(t, "G", fretboard.strings[2].root.String())
		assert.Equal(t, "D", fretboard.strings[3].root.String())
		assert.Equal(t, "A", fretboard.strings[4].root.String())
		assert.Equal(t, "E", fretboard.strings[5].root.String())
	})

	t.Run("use 22 frets as default for zero value Options", func(t *testing.T) {
		fretboard, err := New(Options{})

		assert.NoError(t, err)
		assert.Equal(t, uint(22), fretboard.Frets)
	})

	t.Run("parse guitar strings from specified tuning", func(t *testing.T) {
		tuning, _ := NewTuning("D A D G B E")

		fretboard, err := New(Options{Tuning: tuning})

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
