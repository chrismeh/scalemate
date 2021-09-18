package renderer

import (
	"github.com/chrismeh/scalemate/pkg/fretboard"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/colornames"
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
	font             *truetype.Font
}

func NewPNGRenderer(fretboard *fretboard.Fretboard) PNGRenderer {
	fbOffsetX, fbOffsetY := 40.0, 50.0
	stringSpacing, fretSpacing := 30.0, 60.0

	fbWidth := float64(fretboard.Frets) * fretSpacing
	fbHeight := float64(fretboard.Strings) * stringSpacing
	extraSpaceHeadstock := 30.0

	width := int(2*fbOffsetX + fbWidth + extraSpaceHeadstock)
	height := int(2*fbOffsetY + fbHeight)

	dc := gg.NewContext(width, height)

	return PNGRenderer{
		dc:               dc,
		fb:               fretboard,
		width:            width,
		height:           height,
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
	p.font = f

	err = p.drawFretboard(p.fretboardOffsetX, p.fretboardOffsetY)
	if err != nil {
		return err
	}

	return png.Encode(w, p.dc.Image())
}

func (p PNGRenderer) drawFretboard(x, y float64) error {
	p.fillBackground()

	p.drawTitle()
	p.dc.SetFontFace(truetype.NewFace(p.font, &truetype.Options{Size: 12}))
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

func (p PNGRenderer) drawTitle() {
	p.dc.SetFontFace(truetype.NewFace(p.font, &truetype.Options{Size: 20}))
	p.dc.DrawString(p.fb.String(), p.fretboardOffsetX, 0.75*p.fretboardOffsetY)
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

	headStockOutlineX := offsetX + float64(p.fb.Frets)*p.fretSpacing
	headStockOutlineTopY := offsetY + p.stringSpacing
	headStockOutlineBottomY := offsetY + float64(p.fb.Strings)*p.stringSpacing
	p.dc.DrawLine(headStockOutlineX, headStockOutlineTopY, float64(p.width)-p.fretboardOffsetX, headStockOutlineTopY-20)
	p.dc.DrawLine(headStockOutlineX, headStockOutlineBottomY, float64(p.width)-p.fretboardOffsetX, headStockOutlineBottomY+20)

	p.dc.Stroke()
}

func (p PNGRenderer) drawTuning(offsetX, offsetY float64) {
	notes := p.fb.Tuning.Notes()
	for i := 0; i < len(notes); i++ {
		stringNumber := int(p.fb.Strings) - i
		x := offsetX
		y := offsetY + float64(stringNumber)*p.stringSpacing

		p.drawNote(notes[i], x, y)
	}
}

func (p PNGRenderer) drawHighlightedNotes(offsetX, offsetY float64) error {
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
			p.drawNote(fret.Note, x, y)
		}
	}
	return nil
}

func (p PNGRenderer) drawNote(note string, x, y float64) {
	switch {
	case p.fb.Scale.Root() == note:
		p.dc.SetColor(colornames.Lightblue)
	case p.fb.Scale.Contains(note):
		p.dc.SetColor(colornames.Black)
	default:
		p.dc.SetColor(colornames.Grey)
	}
	p.dc.DrawCircle(x, y, 10)
	p.dc.Fill()

	p.dc.SetColor(colornames.White)
	p.dc.DrawStringAnchored(note, x, y-2, 0.5, 0.5)
	p.dc.Stroke()
}
