package main

var notes = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

type Note struct {
	value string
}

func (n Note) Add(halfsteps uint) Note {
	if halfsteps == 0 || halfsteps%12 == 0 {
		return n
	}

	var currentNoteIndex uint = 0
	for i, note := range notes {
		if n.value == note {
			currentNoteIndex = uint(i)
			break
		}
	}

	nextNoteIndex := currentNoteIndex + halfsteps
	if int(nextNoteIndex) > 11 {
		nextNoteIndex -= 12
	}

	return Note{value: notes[nextNoteIndex]}
}

func (n Note) String() string {
	return n.value
}
