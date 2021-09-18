package renderer

import (
	"github.com/chrismeh/scalemate/pkg/fretboard"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image/png"
	"io"
)

type PNGRenderer struct {
	dc               *gg.Context
	fb               *fretboard.Fretboard
	width            int
	height           int
	fretboardOffsetX float64
	fretboardOffsetY float64
	stringSpacing    float64
	fretSpacing      float64
	fontRegular      font.Face
	fontSmall        font.Face
}

func NewPNGRenderer(fretboard *fretboard.Fretboard) PNGRenderer {
	fbOffsetX, fbOffsetY := 40.0, 50.0
	stringSpacing, fretSpacing := 30.0, 60.0

	fbWidth := float64(fretboard.Frets) * fretSpacing
	fbHeight := float64(fretboard.Strings) * stringSpacing
	extraSpaceTuning := 15.0

	width := 2*fbOffsetX + fbWidth + extraSpaceTuning
	height := 2*fbOffsetY + fbHeight

	dc := gg.NewContext(int(width), int(height))

	return PNGRenderer{
		dc:               dc,
		fb:               fretboard,
		width:            int(width),
		height:           int(height),
		fretboardOffsetX: fbOffsetX,
		fretboardOffsetY: fbOffsetY,
		stringSpacing:    stringSpacing,
		fretSpacing:      fretSpacing,
	}
}

func (p PNGRenderer) Render(w io.Writer) error {
	f, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return err
	}
	p.fontRegular = truetype.NewFace(f, &truetype.Options{Size: 14})
	p.fontSmall = truetype.NewFace(f, &truetype.Options{Size: 12})
	p.fillBackground()

	p.dc.SetFontFace(truetype.NewFace(f, &truetype.Options{Size: 20}))
	p.dc.DrawString(p.fb.String(), p.fretboardOffsetX, 0.75*p.fretboardOffsetY)

	p.dc.SetFontFace(truetype.NewFace(f, &truetype.Options{Size: 14}))
	err = p.drawFretboard(p.fretboardOffsetX, p.fretboardOffsetY)
	if err != nil {
		return err
	}

	return png.Encode(w, p.dc.Image())
}

func (p PNGRenderer) drawFretboard(x, y float64) error {
	p.drawNeck(x, y)
	p.drawTuning(x+float64(p.fb.Frets)*p.fretSpacing, y)

	err := p.drawHighlightedNotes(x, y)
	if err != nil {
		return err
	}

	return nil
}

func (p PNGRenderer) fillBackground() {
	p.dc.SetColor(colornames.White)
	p.dc.DrawRectangle(0, 0, float64(p.width), float64(p.height))
	p.dc.Fill()

	p.dc.SetColor(colornames.Black)
}

func (p PNGRenderer) drawNeck(offsetX, offsetY float64) {
	for str := 1; str <= int(p.fb.Strings); str++ {
		p.dc.DrawLine(
			offsetX,
			offsetY+float64(str)*p.stringSpacing,
			offsetX+float64(p.fb.Frets)*p.fretSpacing,
			offsetY+float64(str)*p.stringSpacing,
		)
	}

	for fret := 0; fret <= int(p.fb.Frets); fret++ {
		p.dc.DrawLine(
			offsetX+float64(fret)*p.fretSpacing,
			offsetY+p.stringSpacing,
			offsetX+float64(fret)*p.fretSpacing,
			offsetY+float64(p.fb.Strings)*p.stringSpacing,
		)
	}

	p.dc.Stroke()
}

func (p PNGRenderer) drawTuning(offsetX, offsetY float64) {
	notes := p.fb.Tuning.Notes()
	for i := 0; i < len(notes); i++ {
		stringNumber := int(p.fb.Strings) - i
		p.dc.DrawString(notes[i], offsetX+10, offsetY+float64(stringNumber)*p.stringSpacing+5)
	}
	p.dc.Stroke()
}

func (p PNGRenderer) drawHighlightedNotes(offsetX, offsetY float64) error {
	p.dc.SetFontFace(p.fontSmall)

	for s := 1; s <= int(p.fb.Strings); s++ {
		for f := int(p.fb.Frets); f > 0; f-- {
			fret, err := p.fb.Fret(uint(s), uint(f))
			if err != nil {
				return err
			}
			if !fret.Highlighted {
				continue
			}

			x := offsetX + float64(p.fb.Frets-uint(f-1))*p.fretSpacing - 0.5*p.fretSpacing
			y := offsetY + float64(s)*p.stringSpacing
			p.dc.SetColor(colornames.Black)
			if fret.Root {
				p.dc.SetColor(colornames.Lightblue)
			}
			p.dc.DrawCircle(x, y, 10)
			p.dc.Fill()

			p.dc.SetColor(colornames.White)
			p.dc.DrawStringAnchored(fret.Note.String(), x, y-2, 0.5, 0.5)
			p.dc.Stroke()
		}
	}
	return nil
}
