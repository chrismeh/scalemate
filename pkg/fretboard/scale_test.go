package fretboard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScale_Chords(t *testing.T) {
	t.Run("create chords for a major scale", func(t *testing.T) {
		scale, _ := NewScale("C", ScaleMajor)
		chords := scale.Chords()

		assert.Equal(t, "Cmaj7", chords[0].Name)
		assert.Equal(t, "Dmin7", chords[1].Name)
		assert.Equal(t, "Emin7", chords[2].Name)
		assert.Equal(t, "Fmaj7", chords[3].Name)
		assert.Equal(t, "G7", chords[4].Name)
		assert.Equal(t, "Amin7", chords[5].Name)
		assert.Equal(t, "Bmin7b5", chords[6].Name)
	})

	t.Run("create chords for a minor scale", func(t *testing.T) {
		scale, _ := NewScale("A", ScaleMinor)
		chords := scale.Chords()

		assert.Equal(t, "Amin7", chords[0].Name)
		assert.Equal(t, "Bmin7b5", chords[1].Name)
		assert.Equal(t, "Cmaj7", chords[2].Name)
		assert.Equal(t, "Dmin7", chords[3].Name)
		assert.Equal(t, "Emin7", chords[4].Name)
		assert.Equal(t, "Fmaj7", chords[5].Name)
		assert.Equal(t, "G7", chords[6].Name)
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

func TestNote_Equals(t *testing.T) {
	t.Run("return true when both notes have the same value", func(t *testing.T) {
		a, b := Note{value: "C"}, Note{value: "C"}

		assert.True(t, a.Equals(b))
	})

	t.Run("return false when notes have different values", func(t *testing.T) {
		a, b := Note{value: "C"}, Note{value: "C#"}

		assert.False(t, a.Equals(b))
	})
}

func TestNote_Add(t *testing.T) {
	tests := []struct {
		Name         string
		StartNote    string
		Halfsteps    uint
		ExpectedNote string
	}{
		{Name: "Add 0 semitones from A", StartNote: "A", Halfsteps: 0, ExpectedNote: "A"},
		{Name: "Add 1 semitone from A", StartNote: "A", Halfsteps: 1, ExpectedNote: "A#"},
		{Name: "Add 2 semitones from A", StartNote: "A", Halfsteps: 2, ExpectedNote: "B"},
		{Name: "Add 12 semitones from A", StartNote: "A", Halfsteps: 12, ExpectedNote: "A"},
		{Name: "Add 13 semitones from A", StartNote: "A", Halfsteps: 13, ExpectedNote: "A#"},
		{Name: "Add 1 semitone from E", StartNote: "E", Halfsteps: 1, ExpectedNote: "F"},
		{Name: "Add 5 semitones from E", StartNote: "E", Halfsteps: 5, ExpectedNote: "A"},
		{Name: "Add 22 semitones from E", StartNote: "E", Halfsteps: 22, ExpectedNote: "D"},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			start := Note{value: tt.StartNote}
			result := start.Add(tt.Halfsteps)

			assert.Equal(t, tt.ExpectedNote, result.String())
		})
	}
}

func TestNote_IntervalTo(t *testing.T) {
	tests := []struct {
		Name             string
		FirstNoteValue   string
		SecondNoteValue  string
		ExpectedInterval uint
	}{
		{Name: "perfect unison for equal notes", FirstNoteValue: "G", SecondNoteValue: "G", ExpectedInterval: 1},
		{Name: "second for a single semitone", FirstNoteValue: "G", SecondNoteValue: "G#", ExpectedInterval: 2},
		{Name: "second for two semitones", FirstNoteValue: "G", SecondNoteValue: "A", ExpectedInterval: 2},
		{Name: "third for three semitones", FirstNoteValue: "G", SecondNoteValue: "A#", ExpectedInterval: 3},
		{Name: "third for four semitones", FirstNoteValue: "G", SecondNoteValue: "B", ExpectedInterval: 3},
		{Name: "perfect fourth for five semitones", FirstNoteValue: "G", SecondNoteValue: "C", ExpectedInterval: 4},
		{Name: "perfect fifth for seven semitones", FirstNoteValue: "G", SecondNoteValue: "D", ExpectedInterval: 5},
		{Name: "sixth for eight semitones", FirstNoteValue: "G", SecondNoteValue: "D#", ExpectedInterval: 6},
		{Name: "sixth for nine semitones", FirstNoteValue: "G", SecondNoteValue: "E", ExpectedInterval: 6},
		{Name: "seventh for ten semitones", FirstNoteValue: "G", SecondNoteValue: "F", ExpectedInterval: 7},
		{Name: "seventh for nine semitones", FirstNoteValue: "G", SecondNoteValue: "F#", ExpectedInterval: 7},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			n1, n2 := Note{value: tt.FirstNoteValue}, Note{value: tt.SecondNoteValue}
			assert.Equal(t, tt.ExpectedInterval, n1.IntervalTo(n2))
		})
	}
}
