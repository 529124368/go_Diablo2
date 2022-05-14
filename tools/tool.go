package tools

import (
	"bytes"
	"container/ring"
	"fmt"
	"game/engine/ws/pb"
	"image"
	_ "image/png"
	"log"
	"math"
	"strings"
	"time"

	"github.com/fzipp/texturepacker"
	"github.com/golang/protobuf/proto"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	IDLE                uint8         = 0
	Walk                uint8         = 1
	RUN                 uint8         = 3
	ATTACK              uint8         = 2
	SkILL               uint8         = 4
	SPEED               float64       = 2 //玩家走路移动速度
	SPEED_RUN           float64       = 3 //玩家跑步移动速度
	ISHIDDEN            uint8         = 1 //装备栏等隐藏标识
	ISITEMS             uint8         = 0 //物品和装备标识
	ISMINICOM           uint8         = 2 //MINI板标识
	ISNORCOM            uint8         = 3 //无标识 占位用
	LAYOUTX             int           = 790
	LAYOUTY             int           = 480
	CLOSEBTNSLEEP       time.Duration = 200000000 //按钮按下弹起动画sleep时间
	MUSICWAV            int           = 1         //音乐WAV格式
	MUSICMP3            int           = 2         //音乐mp3格式
	ObjectPath          string        = "resource/object"
	GAMESCENELOGIN      int           = 1 //登录场景
	GAMESCENESELECTROLE int           = 2 //选择场景
	GAMESCENEOPENDOOR   int           = 3 //开门场景
	GAMESCENESTART      int           = 4 //游戏场景
	BgmMusic            uint8         = 0 //音乐类型 背景音乐
	SceneMusic          uint8         = 1 //音乐类型 背景音乐
	PlistN              uint8         = 1
	PlistR              uint8         = 2
)

//Calculate Direction
func CaluteDir(x, y, x_tar, y_tar int64) uint8 {
	len := Distance(x, y, x_tar, y_tar)
	//TODO  240 is PLAYERCENTERY
	a := Angle(math.Abs(float64(int64(y_tar)-240)), len)

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
	pathList := ring.New(16)
	dirList := []uint8{1, 11, 6, 12, 2, 13, 7, 14, 3, 15, 4, 8, 0, 9, 5, 10}
	path1 := make([]uint8, 0)
	path2 := make([]uint8, 0)
	for i := 0; i < 16; i++ {
		if pathList.Value == nil {
			pathList.Value = dirList[i]
			pathList = pathList.Next()
		}
	}
	var oldDir_k, newDir_k int
	for k, v := range dirList {
		if v == oldDir {
			oldDir_k = k
		}
		if v == newDir {
			newDir_k = k
		}
	}
	pathList = pathList.Move(oldDir_k)
	for pathList.Value != newDir {
		pathList = pathList.Next()
		path1 = append(path1, pathList.Value.(uint8))
	}
	pathList = pathList.Move(16 - newDir_k + oldDir_k)
	for pathList.Value != newDir {
		pathList = pathList.Prev()
		path2 = append(path2, pathList.Value.(uint8))
	}
	if len(path1) < len(path2) {
		return path1[:len(path1)-1]
	} else {
		return path2[:len(path2)-1]
	}
}

//获取物品的尺寸
func GetItemsCellSize(name string) (int, int) {
	type1 := "HP0,MP0,neck,ring"
	type2 := "book"
	type3 := "dun,head-4,head-5,hand,shose,head-3,box"
	type4 := "dun-6,sword,sword-1,body-3,body-2,dun-4,dun-5,futou,futou-1,body-4,dun-3,futou-3"
	type5 := "blet"
	if strings.Contains(type1, name) {
		return 1, 1
	} else if strings.Contains(type2, name) {
		return 1, 2
	} else if strings.Contains(type3, name) {
		return 2, 2
	} else if strings.Contains(type4, name) {
		return 2, 3
	} else if strings.Contains(type5, name) {
		return 2, 1
	} else {
		return 2, 4
	}
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
	xx := math.Floor((M_Minus_N + M_Plus_N) / 2)
	yy := math.Floor((M_Minus_N+M_Plus_N)/2 - M_Minus_N)
	return int(xx), int(yy)
}

//根据方向计算偏移距离
func CalculateSpeed(dir uint8, speed float64) (float64, float64) {
	moveX, moveY := 0.0, 0.0
	switch dir {
	case 0:
		moveX, moveY = -speed, speed
	case 1:
		moveX, moveY = -speed, -speed
	case 2:
		moveX, moveY = speed, -speed
	case 3:
		moveX, moveY = speed, speed
	case 4:
		moveX, moveY = 0, speed
	case 5:
		moveX, moveY = -speed, 0
	case 6:
		moveX, moveY = 0, -speed
	case 7:
		moveX, moveY = speed, 0
	case 8:
		moveX, moveY = 1-speed, speed
	case 9:
		moveX, moveY = -speed, speed-1
	case 10:
		moveX, moveY = -speed, 1-speed
	case 11:
		moveX, moveY = 1-speed, -speed
	case 12:
		moveX, moveY = speed-1, -speed
	case 13:
		moveX, moveY = speed, 1-speed
	case 14:
		moveX, moveY = speed, speed-1
	case 15:
		moveX, moveY = speed-1, speed
	}
	return moveX, moveY
}

type OffsetXY struct {
	X, Y int
}

//根据玩家动作和加载的资源获取偏移
func GetOffetByAction(name string) [4]OffsetXY {

	var box [4]OffsetXY
	switch name {
	case "ba":
		box[0] = OffsetXY{-4, -3}
		box[1] = OffsetXY{-4, -3}
		box[2] = OffsetXY{-50, -30}
		box[3] = OffsetXY{-50, -30}
	case "ba1":
		box[0] = OffsetXY{4, -18}
		box[1] = OffsetXY{4, -18}
		box[2] = OffsetXY{-55, -35}
		box[3] = OffsetXY{-55, -35}
	case "ba2":
		box[0] = OffsetXY{3, -7}
		box[1] = OffsetXY{8, -10}
		box[2] = OffsetXY{-55, -20}
		box[3] = OffsetXY{-10, -15}
	}
	return box
}

//解包
func Unpack(msg []byte) *pb.Message {
	m := &pb.Message{}
	err := proto.Unmarshal(msg, m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}

//消息打包
func Pack(s bool, f, datas, msg string, p *pb.Player) []byte {
	m := &pb.Message{
		Flag: f,
		Data: &pb.Datas{
			Status: s,
			Data:   datas,
			Mes:    msg,
			Man:    p,
		},
	}
	d, err := proto.Marshal(m)
	if err != nil {
		fmt.Println("has error###", err)
	}
	return d
}
