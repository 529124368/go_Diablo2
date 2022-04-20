package fonts

import (
	"embed"
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type FontBase struct {
	Size   float64
	DPI    float64
	tt     *truetype.Font
	asset  *embed.FS
	f_face font.Face
}

func NewFont(ass *embed.FS) *FontBase {
	f := &FontBase{
		Size:   10,
		DPI:    150,
		tt:     nil,
		f_face: nil,
		asset:  ass,
	}
	return f
}

func (f *FontBase) LoadFont(path string) {
	ff, err := f.asset.ReadFile(path)
	if err != nil {
		log.Fatal("font path error")
	}
	tt, _ := truetype.Parse(ff)
	if err != nil {
		log.Fatal(err)
	}
	f.tt = tt
	face := truetype.NewFace(f.tt, &truetype.Options{
		Size:    f.Size,
		DPI:     f.DPI,
		Hinting: font.HintingFull,
	})
	f.f_face = face
}

func (f *FontBase) RenderByCustomer(screen *ebiten.Image, size, dpi float64, x, y int, cont string) {
	f_face := truetype.NewFace(f.tt, &truetype.Options{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	text.Draw(screen, cont, f_face, x, y, color.RGBA{R: 255, G: 0, B: 0, A: 255})
}

func (f *FontBase) Render(screen *ebiten.Image, x, y int, cont string, size, dpi float64, f_color color.Color) {
	if f.Size != size || f.DPI != dpi {
		f.f_face.Close()
		face := truetype.NewFace(f.tt, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: font.HintingFull,
		})
		f.Size = size
		f.DPI = dpi
		f.f_face = face
	}
	text.Draw(screen, cont, f.f_face, x, y, f_color)
}
