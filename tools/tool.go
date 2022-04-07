package tools

import (
	"bytes"
	"image"
	"log"
	"math"
	"time"

	"github.com/fzipp/texturepacker"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	IDLE          uint8         = 0
	RUN           uint8         = 1
	ATTACK        uint8         = 2
	SPEED         float64       = 2
	ISHIDDENCOM   uint8         = 1
	ISMINICOM     uint8         = 2
	ISNORCOM      uint8         = 3
	LAYOUTX       int           = 790
	LAYOUTY       int           = 480
	CLOSEBTNSLEEP time.Duration = 200000000
)

//Calculate Direction
func CaluteDir(x, y, x_tar, y_tar int64) uint8 {
	len := Distance(x, y, x_tar, y_tar)
	//TODO  242 is PLAYERCENTERY
	a := Angle(math.Abs(float64(int64(y_tar)-242)), len)

	if x < x_tar && y > y_tar {
		if a > 0 && a <= 30 {
			return 13
		}
		if a > 30 && a <= 60 {
			return 2
		}
		if a > 60 && a < 90 {
			return 12
		}
	}
	if x < x_tar && y < y_tar {
		if a > 0 && a <= 30 {
			return 14
		}
		if a > 30 && a <= 60 {
			return 3
		}
		if a > 60 && a < 90 {
			return 15
		}
	}

	if x > x_tar && y < y_tar {
		if a > 0 && a <= 30 {
			return 9
		}
		if a > 30 && a <= 60 {
			return 0
		}
		if a > 60 && a < 90 {
			return 8
		}
	}
	if x > x_tar && y > y_tar {
		if a > 0 && a <= 30 {
			return 10
		}
		if a > 30 && a <= 60 {
			return 1
		}
		if a > 60 && a < 90 {
			return 11
		}
	}

	if x > x_tar && float64(y) == math.Abs(float64(y_tar)) {
		return 5
	}
	if x < x_tar && float64(y) == math.Abs(float64(y_tar)) {
		return 7
	}
	if float64(x) == math.Abs(float64(x_tar)) && y > y_tar {
		return 6
	}
	if float64(x) == math.Abs(float64(x_tar)) && y < y_tar {
		return 4
	}
	return 0
}

//Get Images From Byte
func GetEbitenImage(data []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

//Get NRGBA Plist Images
func GetImageFromPlist(s []byte, json []byte) (*texturepacker.SpriteSheet, *image.NRGBA) {
	sheet, err := texturepacker.SheetFromData(json, texturepacker.FormatJSONHash{})

	img, _, err := image.Decode(bytes.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	sheetImage := img.(*image.NRGBA)
	return sheet, sheetImage
}

//Get Paletted Plist Images
func GetImageFromPlistPaletted(s []byte, json []byte) (*texturepacker.SpriteSheet, *image.Paletted) {
	sheet, err := texturepacker.SheetFromData(json, texturepacker.FormatJSONHash{})

	img, _, err := image.Decode(bytes.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	sheetImage := img.(*image.Paletted)
	return sheet, sheetImage
}

//Calculate Distance
func Distance(xa, ya, xb, yb int64) float64 {
	x := math.Abs(float64(xa - xb))
	y := math.Abs(float64(ya - yb))
	return math.Sqrt(x*x + y*y)
}

//Calculate Angle
func Angle(y float64, len float64) float64 {
	return math.Asin(y/len) * 180 / math.Pi
}
