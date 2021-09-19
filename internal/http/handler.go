package http

import (
	"github.com/chrismeh/scalemate/pkg/fretboard"
	"github.com/chrismeh/scalemate/pkg/renderer"
	"net/http"
)

func (a Application) handleGetScale(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fb, err := fretboard.New(fretboard.Options{Frets: 12})
	if err != nil {
		a.errorLog.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	scale, err := fretboard.NewScale("A", fretboard.ScaleMinor)
	if err != nil {
		a.errorLog.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fb.HighlightScale(scale)

	png := renderer.NewPNGRenderer(fb)

	w.Header().Add("content-type", "image/png")
	err = png.Render(w)
	if err != nil {
		a.errorLog.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
