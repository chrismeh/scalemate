package main

type Fretboard struct {
	tuning  string
	strings []guitarString
	scale   Scale
}

func NewFretboard(tuning string, scale Scale) (Fretboard, error) {
	strings, err := buildStringsFromTuning(tuning)
	if err != nil {
		return Fretboard{}, err
	}

	f := Fretboard{tuning: tuning, strings: strings, scale: scale}
	return f, nil
}

type guitarString struct {
	Root   Note
	Number uint
}

func buildStringsFromTuning(tuning string) ([]guitarString, error) {
	strings := make([]guitarString, len(tuning))
	for i := 0; i < len(tuning); i++ {
		rootNote, err := NewNote(string(tuning[i]))
		if err != nil {
			return nil, err
		}

		stringNumber := len(tuning) - i
		strings[stringNumber-1] = guitarString{
			Root:   rootNote,
			Number: uint(stringNumber),
		}
	}

	return strings, nil
}
