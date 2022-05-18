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
	f_face [3]font.Face
}

func NewFont(ass *embed.FS) *FontBase {
	f := &FontBase{
		Size:  10,
		DPI:   150,
		tt:    nil,
		asset: ass,
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
	//type 1
	face := truetype.NewFace(f.tt, &truetype.Options{
		Size:    8,
		DPI:     130,
		Hinting: font.HintingFull,
	})
	f.f_face[0] = face
	//type 2
	face = truetype.NewFace(f.tt, &truetype.Options{
		Size:    7.2,
		DPI:     150,
		Hinting: font.HintingFull,
	})
	f.f_face[1] = face
	//type 2
	face = truetype.NewFace(f.tt, &truetype.Options{
		Size:    7.2,
		DPI:     120,
		Hinting: font.HintingFull,
	})
	f.f_face[2] = face
}

func (f *FontBase) Render(screen *ebiten.Image, i uint8, x, y int, cont string, size, dpi float64, f_color color.Color) {
	text.Draw(screen, cont, f.f_face[i], x, y, f_color)
}
