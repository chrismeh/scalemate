package http

import (
	"github.com/chrismeh/scalemate/pkg/fretboard"
	"github.com/chrismeh/scalemate/pkg/renderer"
	"io"
	"net/http"
)

func (a Application) handleGetIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	template, err := a.templateFS.Open("index.html")
	if err != nil {
		a.internalServerError(err, w)
		return
	}
	defer template.Close()

	index, err := io.ReadAll(template)
	if err != nil {
		a.internalServerError(err, w)
		return
	}

	_, _ = w.Write(index)
}

func (a Application) handleGetScale(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fb, err := fretboard.New(fretboard.Options{Frets: 12})
	if err != nil {
		a.internalServerError(err, w)
		return
	}

	scale, err := fretboard.NewScale("A", fretboard.ScaleMinor)
	if err != nil {
		a.internalServerError(err, w)
		return
	}
	fb.HighlightScale(scale)

	png := renderer.NewPNGRenderer(fb)

	w.Header().Add("content-type", "image/png")
	err = png.Render(w)
	if err != nil {
		a.internalServerError(err, w)
		return
	}
}
