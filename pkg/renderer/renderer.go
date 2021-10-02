package renderer

import (
	"github.com/chrismeh/scalemate/pkg/fretboard"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/gofont/goregular"
	"image/color"
	"image/png"
	"io"
	"strconv"
)

var (
	colorRootNote  = color.RGBA{R: 0x00, G: 0xd1, B: 0xb2, A: 0xff}
	colorChordNote = color.RGBA{R: 0x98, G: 0x36, B: 0x28, A: 0xff}
	colorScaleNote = color.RGBA{R: 0x08, G: 0x09, B: 0x0a, A: 0xff}
	colorMiscNote  = color.RGBA{R: 0xa4, G: 0x96, B: 0x9b, A: 0xff}
)

type PNGRenderer struct {
	dc            *gg.Context
	fb            *fretboard.Fretboard
	width         int
	height        int
	stringSpacing float64
	fretSpacing   float64
	font          *truetype.Font
	options       PNGOptions
}

type PNGOptions struct {
	FretboardOffsetX float64
	FretboardOffsetY float64
	DrawTitle        bool
}

func NewPNGRenderer(fretboard *fretboard.Fretboard, options PNGOptions) PNGRenderer {
	stringSpacing, fretSpacing := 30.0, 60.0

	fbWidth := float64(fretboard.Frets) * fretSpacing
	fbHeight := float64(fretboard.Strings) * stringSpacing
	extraSpaceHeadstock := 30.0

	width := int(2*options.FretboardOffsetX + fbWidth + extraSpaceHeadstock)
	height := int(2*options.FretboardOffsetY + fbHeight)

	dc := gg.NewContext(width, height)

	return PNGRenderer{
		dc:            dc,
		fb:            fretboard,
		width:         width,
		height:        height,
		stringSpacing: stringSpacing,
		fretSpacing:   fretSpacing,
		options:       options,
	}
}

func (p PNGRenderer) Render(w io.Writer) error {
	f, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return err
	}
	p.font = f

	err = p.drawFretboard()
	if err != nil {
		return err
	}

	return png.Encode(w, p.dc.Image())
}

func (p PNGRenderer) drawFretboard() error {
	p.fillBackground()

	if p.options.DrawTitle {
		p.drawTitle()
	}
	p.drawNeck()
	p.drawTuning()

	err := p.drawHighlightedNotes()
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
	p.dc.DrawString(p.fb.String(), p.options.FretboardOffsetX, 0.75*p.options.FretboardOffsetY)
	p.dc.SetFontFace(truetype.NewFace(p.font, &truetype.Options{Size: 12}))
}

func (p PNGRenderer) drawNeck() {
	for str := 1; str <= int(p.fb.Strings); str++ {
		p.dc.DrawLine(
			p.options.FretboardOffsetX,
			p.options.FretboardOffsetY+float64(str)*p.stringSpacing,
			p.options.FretboardOffsetX+float64(p.fb.Frets)*p.fretSpacing,
			p.options.FretboardOffsetY+float64(str)*p.stringSpacing,
		)
	}

	for fret := 0; fret <= int(p.fb.Frets); fret++ {
		p.dc.DrawLine(
			p.options.FretboardOffsetX+float64(fret)*p.fretSpacing,
			p.options.FretboardOffsetY+p.stringSpacing,
			p.options.FretboardOffsetX+float64(fret)*p.fretSpacing,
			p.options.FretboardOffsetY+float64(p.fb.Strings)*p.stringSpacing,
		)

		fretNumber := int(p.fb.Frets) - fret
		if fretNumber > 0 {
			p.dc.DrawStringAnchored(
				strconv.Itoa(fretNumber),
				p.options.FretboardOffsetX+float64(fret)*p.fretSpacing+0.5*p.fretSpacing,
				p.options.FretboardOffsetY+float64(p.fb.Strings)*p.stringSpacing+0.75*p.stringSpacing,
				0.5,
				0.5,
			)
		}
	}

	headStockOutlineX := p.options.FretboardOffsetX + float64(p.fb.Frets)*p.fretSpacing
	headStockOutlineTopY := p.options.FretboardOffsetY + p.stringSpacing
	headStockOutlineBottomY := p.options.FretboardOffsetY + float64(p.fb.Strings)*p.stringSpacing
	p.dc.DrawLine(headStockOutlineX, headStockOutlineTopY, float64(p.width)-p.options.FretboardOffsetX, headStockOutlineTopY-20)
	p.dc.DrawLine(headStockOutlineX, headStockOutlineBottomY, float64(p.width)-p.options.FretboardOffsetX, headStockOutlineBottomY+20)

	p.dc.Stroke()
}

func (p PNGRenderer) drawTuning() {
	notes := p.fb.Tuning.Notes()
	for i := 0; i < len(notes); i++ {
		stringNumber := int(p.fb.Strings) - i
		x := p.options.FretboardOffsetX + float64(p.fb.Frets)*p.fretSpacing
		y := p.options.FretboardOffsetY + float64(stringNumber)*p.stringSpacing

		p.drawNote(notes[i], x, y)
	}
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

			x := p.options.FretboardOffsetX + float64(p.fb.Frets-uint(f-1))*p.fretSpacing - 0.5*p.fretSpacing
			y := p.options.FretboardOffsetY + float64(s)*p.stringSpacing
			p.drawNote(fret.Note, x, y)
		}
	}
	return nil
}

func (p PNGRenderer) drawNote(note fretboard.Note, x, y float64) {
	switch {
	case p.fb.Scale.Root == note:
		p.dc.SetColor(colorRootNote)
	case p.fb.Chord.Contains(note):
		p.dc.SetColor(colorChordNote)
	case p.fb.Scale.Contains(note):
		p.dc.SetColor(colorScaleNote)
	default:
		p.dc.SetColor(colorMiscNote)
	}
	p.dc.DrawCircle(x, y, 10)
	p.dc.Fill()

	p.dc.SetColor(colornames.White)
	p.dc.DrawStringAnchored(note.String(), x, y-2, 0.5, 0.5)
	p.dc.Stroke()
}
