package fretboard

type Chord struct {
	root      Note
	intervals []uint
}

func (c Chord) Notes() []Note {
	notes := make([]Note, len(c.intervals)+1)
	notes[0] = c.root

	for i, v := range c.intervals {
		notes[i+1] = c.root.Add(v)
	}

	return notes
}
