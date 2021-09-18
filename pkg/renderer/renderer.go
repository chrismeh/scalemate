package renderer

import (
	"github.com/chrismeh/scalemate/pkg/fretboard"
	"github.com/fogleman/gg"
	"golang.org/x/image/colornames"
	"image/png"
	"io"
)

type PNGRenderer struct {
	dc            *gg.Context
	fb            *fretboard.Fretboard
	width         int
	height        int
	margin        float64
	stringSpacing float64
	fretSpacing   float64
	stringsX1     float64
	stringsX2     float64
	fretsY1       float64
	fretsY2       float64
}

func NewPNGRenderer(fretboard *fretboard.Fretboard) PNGRenderer {
	margin, stringSpacing, fretSpacing := 40.0, 30.0, 60.0
	width := 2*margin + float64(fretboard.Frets)*fretSpacing
	height := 2*margin + float64(fretboard.Strings)*stringSpacing
	dc := gg.NewContext(int(width), int(height))

	return PNGRenderer{
		dc:            dc,
		fb:            fretboard,
		width:         int(width),
		height:        int(height),
		margin:        margin,
		stringSpacing: stringSpacing,
		fretSpacing:   fretSpacing,
		stringsX1:     margin,
		stringsX2:     width - margin,
		fretsY1:       stringSpacing,
		fretsY2:       float64(fretboard.Strings) * stringSpacing,
	}
}

func (p PNGRenderer) Render(w io.Writer) error {
	p.drawNeck()
	p.drawTuning()

	err := p.drawHighlightedNotes()
	if err != nil {
		return err
	}

	return png.Encode(w, p.dc.Image())
}

func (p PNGRenderer) drawNeck() {
	p.dc.SetColor(colornames.White)
	p.dc.DrawRectangle(0, 0, float64(p.width), float64(p.height))
	p.dc.Fill()

	p.dc.SetColor(colornames.Black)
	for x := 1; x <= int(p.fb.Strings); x++ {
		p.dc.DrawLine(
			p.stringsX1,
			float64(x)*p.stringSpacing,
			p.stringsX2,
			float64(x)*p.stringSpacing,
		)
	}

	for x := 0; x <= int(p.fb.Frets); x++ {
		p.dc.DrawLine(
			p.margin+float64(x)*p.fretSpacing,
			p.fretsY1,
			p.margin+float64(x)*p.fretSpacing,
			p.fretsY2,
		)
	}

	p.dc.Stroke()
}

func (p PNGRenderer) drawTuning() {
	notes := p.fb.Tuning.Notes()
	for i := 0; i < len(notes); i++ {
		stringNumber := int(p.fb.Strings) - i
		p.dc.DrawString(notes[i], p.stringsX2+10, float64(stringNumber)*p.stringSpacing+5)
	}
	p.dc.Stroke()
}

func (p PNGRenderer) drawHighlightedNotes() error {
	for s := 1; s <= int(p.fb.Strings); s++ {
		for f := int(p.fb.Frets); f > 0; f-- {
			fret, err := p.fb.Fret(uint(s), uint(f))
			if err != nil {
				return err
			}
			if !fret.Highlighted {
				continue
			}

			x := float64(p.fb.Frets-uint(f-1))*p.fretSpacing + p.margin - 0.5*p.fretSpacing
			p.dc.SetColor(colornames.Black)
			if fret.Root {
				p.dc.SetColor(colornames.Lightblue)
			}
			p.dc.DrawCircle(x, float64(s)*p.stringSpacing, 10)
			p.dc.Fill()
		}
	}
	return nil
}
