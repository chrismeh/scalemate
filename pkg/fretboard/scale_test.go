package fretboard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScale_Chords(t *testing.T) {
	t.Run("create chords for a major scale", func(t *testing.T) {
		scale, _ := NewScale("C", ScaleMajor)
		chords := scale.Chords()

		assert.Equal(t, "Cmaj7", chords[0].Name())
		assert.Equal(t, "Dmin7", chords[1].Name())
		assert.Equal(t, "Emin7", chords[2].Name())
		assert.Equal(t, "Fmaj7", chords[3].Name())
		assert.Equal(t, "G7", chords[4].Name())
		assert.Equal(t, "Amin7", chords[5].Name())
		assert.Equal(t, "Bmin7b5", chords[6].Name())
	})

	t.Run("create chords for a minor scale", func(t *testing.T) {
		scale, _ := NewScale("A", ScaleMinor)
		chords := scale.Chords()

		assert.Equal(t, "Amin7", chords[0].Name())
		assert.Equal(t, "Bmin7b5", chords[1].Name())
		assert.Equal(t, "Cmaj7", chords[2].Name())
		assert.Equal(t, "Dmin7", chords[3].Name())
		assert.Equal(t, "Emin7", chords[4].Name())
		assert.Equal(t, "Fmaj7", chords[5].Name())
		assert.Equal(t, "G7", chords[6].Name())
	})
}

func TestScale_Root(t *testing.T) {
	scale, _ := NewScale("A", ScaleMinor)
	assert.Equal(t, Note{value: "A"}, scale.Root)
}

func TestScale_Contains(t *testing.T) {
	t.Run("return true if scale contains note", func(t *testing.T) {
		scale, _ := NewScale("A", ScaleMinor)
		assert.True(t, scale.Contains(Note{value: "C"}))
	})

	t.Run("return false if scale does not contain note", func(t *testing.T) {
		scale, _ := NewScale("A", ScaleMinor)
		assert.False(t, scale.Contains(Note{value: "C#"}))
	})
}

func TestScale_Title(t *testing.T) {
	scale, _ := NewScale("A", ScaleMinor)

	assert.Equal(t, "A minor", scale.Name())
}

func TestNewScale(t *testing.T) {
	t.Run("return error when note does not exist", func(t *testing.T) {
		_, err := NewScale("M", ScaleMinor)
		assert.Error(t, err)
	})

	t.Run("return error when scale type is unknown", func(t *testing.T) {
		_, err := NewScale("A", "Foo")
		assert.Error(t, err)
	})

	t.Run("build correct natural minor scale", func(t *testing.T) {
		testScale, err := NewScale("A", ScaleMinor)
		expectedScale := Scale{
			Root:      Note{value: "A"},
			notes:     []Note{{value: "A"}, {value: "B"}, {value: "C"}, {value: "D"}, {value: "E"}, {value: "F"}, {value: "G"}},
			scaleType: ScaleMinor,
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedScale, testScale)
	})

	t.Run("build correct major scale", func(t *testing.T) {
		testScale, err := NewScale("C", ScaleMajor)
		expectedScale := Scale{
			Root:      Note{value: "C"},
			notes:     []Note{{value: "C"}, {value: "D"}, {value: "E"}, {value: "F"}, {value: "G"}, {value: "A"}, {value: "B"}},
			scaleType: ScaleMajor,
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedScale, testScale)
	})

	t.Run("build correct harmonic minor scale", func(t *testing.T) {
		testScale, err := NewScale("A", ScaleHarmonicMinor)
		expectedScale := Scale{
			Root:      Note{value: "A"},
			notes:     []Note{{value: "A"}, {value: "B"}, {value: "C"}, {value: "D"}, {value: "E"}, {value: "F"}, {value: "G#"}},
			scaleType: ScaleHarmonicMinor,
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedScale, testScale)
	})
}

func TestNewNote(t *testing.T) {
	t.Run("return note struct with correct value", func(t *testing.T) {
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
			start := Note{value: tt.StartNote}
			result := start.Add(tt.Halfsteps)

			assert.Equal(t, tt.ExpectedNote, result.String())
		})
	}
}
