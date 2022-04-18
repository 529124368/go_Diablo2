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
	ISHIDDEN      uint8         = 1 //装备栏等隐藏标识
	ISITEMS       uint8         = 0 //物品和装备标识
	ISMINICOM     uint8         = 2 //MINI板标识
	ISNORCOM      uint8         = 3
	LAYOUTX       int           = 790
	LAYOUTY       int           = 480
	CLOSEBTNSLEEP time.Duration = 200000000 //按钮按下弹起动画sleep时间
	MUSICWAV      int           = 1         //音乐WAV格式
	MUSICMP3      int           = 2         //音乐mp3格式
	ObjectPath    string        = "resource/object"
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
	sheet, _ := texturepacker.SheetFromData(json, texturepacker.FormatJSONHash{})

	img, _, _ := image.Decode(bytes.NewReader(s))
	sheetImage := img.(*image.NRGBA)
	return sheet, sheetImage
}

//Get Paletted Plist Images
func GetImageFromPlistPaletted(s []byte, json []byte) (*texturepacker.SpriteSheet, *image.Paletted) {
	sheet, _ := texturepacker.SheetFromData(json, texturepacker.FormatJSONHash{})

	img, _, _ := image.Decode(bytes.NewReader(s))
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

//计算转头角度一栏
func CalculateDirPath(oldDir, newDir uint8) []uint8 {
	newPath := make([]uint8, 0, 16)
	dirList := []uint8{1, 11, 6, 12, 2, 13, 7, 14, 3, 15, 4, 8, 0, 9, 5, 10}
	newDirIndex := 17
	oldDirIndex := 17
	for k, v := range dirList {
		if v == newDir {
			newDirIndex = k
		}
		if v == oldDir {
			oldDirIndex = k
		}
	}
	if math.Abs(float64(newDirIndex-oldDirIndex)) < 16-math.Abs(float64(newDirIndex-oldDirIndex)) {
		if oldDirIndex < newDirIndex {
			for i := oldDirIndex; i <= newDirIndex; i++ {
				newPath = append(newPath, dirList[i])
			}
		} else {
			for i := oldDirIndex; i >= newDirIndex; i-- {

				newPath = append(newPath, dirList[i])
			}
		}

	} else {
		if oldDirIndex < newDirIndex {
			if oldDirIndex == 0 {
				newPath = append(newPath, dirList[oldDirIndex])
				i := 15
				for i >= newDirIndex {
					newPath = append(newPath, dirList[i])
					i--
				}
			} else {
				i := oldDirIndex
				for i >= 0 {
					newPath = append(newPath, dirList[i])
					i--
				}
				j := 15
				for j >= newDirIndex {
					newPath = append(newPath, dirList[j])
					j--
				}
			}
		} else {
			i := oldDirIndex
			for i <= 15 {
				newPath = append(newPath, dirList[i])
				i++
			}
			j := 0
			for j <= newDirIndex {
				newPath = append(newPath, dirList[j])
				j++
			}
		}

	}
	return newPath[1 : len(newPath)-1]
}

//获取物品的尺寸
func GetItemsCellSize(name string) (int, int) {
	switch name {
	case "HP0":
		return 1, 1
	case "book":
		return 1, 2
	case "dun":
		return 2, 2
	case "head-4":
		return 2, 2
	case "head-5":
		return 2, 2
	case "dun-6":
		return 2, 3
	case "sword":
		return 2, 3
	case "body-3":
		return 2, 3
	case "dun-4":
		return 2, 3
	case "dun-5":
		return 2, 3
	case "hand":
		return 2, 2
	case "shose":
		return 2, 2
	case "head-3":
		return 2, 2
	}

	return 0, 0
}

// AbsInt32 returns the absolute of the given int32
func AbsInt32(a int32) int32 {
	if a < 0 {
		return -a
	}

	return a
}

// MinInt32 returns the higher of two values
func MinInt32(a, b int32) int32 {
	if a < b {
		return a
	}

	return b
}
func MaxInt32(a, b int32) int32 {
	if a > b {
		return a
	}

	return b
}

//计算玩家在所在地砖的逻辑坐标
func GetFloorPositionAt(x, y float64) (int, int) {
	//当前菱形地图 0,0 点坐标世界坐标是 （3280,0）
	M_Minus_N := (x - 3280) / 80
	M_Plus_N := y / 40
	xx := math.Floor((M_Minus_N+M_Plus_N)/2 + 0.5)
	yy := math.Floor(xx - M_Minus_N + 0.5)
	return int(xx), int(yy)
}
