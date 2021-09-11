package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNote_Add(t *testing.T) {
	tests := []struct {
		Name         string
		StartNote    string
		Halfsteps    uint
		ExpectedNote string
	}{
		{Name: "add 0 halfsteps from A", StartNote: "A", Halfsteps: 0, ExpectedNote: "A"},
		{Name: "add 1 halfstep from A", StartNote: "A", Halfsteps: 1, ExpectedNote: "A#"},
		{Name: "add 2 halfsteps from A", StartNote: "A", Halfsteps: 2, ExpectedNote: "B"},
		{Name: "add 12 halfsteps from A", StartNote: "A", Halfsteps: 12, ExpectedNote: "A"},
		{Name: "add 13 halfsteps from A", StartNote: "A", Halfsteps: 13, ExpectedNote: "A#"},
		{Name: "add 1 halfstep from E", StartNote: "E", Halfsteps: 1, ExpectedNote: "F"},
		{Name: "add 5 halfsteps from E", StartNote: "E", Halfsteps: 5, ExpectedNote: "A"},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			start := Note{value: tt.StartNote}
			result := start.Add(tt.Halfsteps)

			assert.Equal(t, tt.ExpectedNote, result.String())
		})
	}
}
