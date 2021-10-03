package http

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/chrismeh/scalemate/pkg/fretboard"
	"github.com/chrismeh/scalemate/pkg/renderer"
	"io"
	"net/http"
	"strconv"
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

	request := parseGetScaleRequest(r)
	fb, err := buildFretboard(request)

	chords := make([]string, 0, 8)
	for _, c := range fb.Scale.Chords() {
		chords = append(chords, c.Name)
	}

	options := renderer.PNGOptions{
		FretboardOffsetX: 0,
		FretboardOffsetY: 40.0,
		DrawTitle:        false,
		TextDisplayMode:  request.displayMode,
	}
	png := renderer.NewPNGRenderer(fb, options)

	var buf bytes.Buffer
	err = png.Render(&buf)
	if err != nil {
		a.internalServerError(err, w)
		return
	}

	resp := struct {
		Picture string   `json:"picture"`
		Chords  []string `json:"chords"`
	}{
		Picture: base64.StdEncoding.EncodeToString(buf.Bytes()),
		Chords:  chords,
	}

	w.Header().Add("content-type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		a.internalServerError(err, w)
		return
	}
}

type getScaleRequest struct {
	rootNote    string
	scaleType   string
	tuning      string
	frets       uint
	chord       string
	displayMode renderer.TextDisplayMode
}

func parseGetScaleRequest(r *http.Request) getScaleRequest {
	req := getScaleRequest{
		rootNote:    "A",
		scaleType:   fretboard.ScaleMinor,
		tuning:      fretboard.TuningStandard,
		frets:       12,
		chord:       "",
		displayMode: renderer.TextDisplayModeDefault,
	}

	query := r.URL.Query()
	if rootNote := query.Get("root"); rootNote != "" {
		req.rootNote = rootNote
	}
	if scaleType := query.Get("type"); scaleType != "" {
		req.scaleType = scaleType
	}
	if tuning := query.Get("tuning"); tuning != "" {
		req.tuning = tuning
	}
	if frets := query.Get("frets"); frets != "" {
		numberOfFrets, err := strconv.Atoi(frets)
		if err == nil && numberOfFrets > 0 {
			req.frets = uint(numberOfFrets)
		}
	}
	if chord := query.Get("chord"); chord != "" {
		req.chord = chord
	}
	if display := query.Get("displayMode"); display != "" {
		displayMode, err := strconv.Atoi(display)
		if err == nil && displayMode > 0 {
			req.displayMode = renderer.TextDisplayMode(displayMode)
		}
	}

	return req
}

func buildFretboard(request getScaleRequest) (*fretboard.Fretboard, error) {
	tuning, err := fretboard.NewTuning(request.tuning)
	if err != nil {
		return nil, err
	}

	fb, err := fretboard.New(fretboard.Options{Frets: request.frets, Tuning: tuning})
	if err != nil {
		return nil, err
	}

	scale, err := fretboard.NewScale(request.rootNote, request.scaleType)
	if err != nil {
		return nil, err
	}
	fb.HighlightScale(scale)

	if request.chord != "" {
		chord, err := fretboard.ParseChord(request.chord)
		if err != nil {
			return nil, err
		}

		fb.HighlightChord(chord)
	}

	return fb, nil
}
