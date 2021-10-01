package fretboard

import "fmt"

var (
	intervalsMajor7          = []uint{4, 7, 11}
	intervalsMinor7          = []uint{3, 7, 10}
	intervalsDominant7       = []uint{4, 7, 10}
	intervalsHalfDiminished7 = []uint{3, 6, 10}
)

type Chord struct {
	root      Note
	intervals []uint
	notes     []Note
}

func (c *Chord) Contains(n Note) bool {
	if len(c.notes) == 0 {
		c.notes = c.buildNotes()
	}

	for _, note := range c.notes {
		if note == n {
			return true
		}
	}

	return false
}

func (c *Chord) buildNotes() []Note {
	notes := make([]Note, len(c.intervals)+1)
	notes[0] = c.root

	for i, v := range c.intervals {
		notes[i+1] = c.root.Add(v)
	}

	return notes
}

func (c *Chord) Name() string {
	var suffix string
	switch {
	case c.compareIntervals(intervalsMajor7...):
		suffix = "maj7"
	case c.compareIntervals(intervalsMinor7...):
		suffix = "min7"
	case c.compareIntervals(intervalsDominant7...):
		suffix = "7"
	case c.compareIntervals(intervalsHalfDiminished7...):
		suffix = "min7b5"
	default:
		suffix = ""
	}

	return fmt.Sprintf("%s%s", c.root, suffix)
}

func (c *Chord) compareIntervals(intervals ...uint) bool {
	if len(c.intervals) != len(intervals) {
		return false
	}

	for i := range c.intervals {
		if c.intervals[i] != intervals[i] {
			return false
		}
	}

	return true
}
