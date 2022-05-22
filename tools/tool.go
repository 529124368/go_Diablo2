package tools

import (
	"bytes"
	"container/ring"
	"fmt"
	"game/engine/ws/pb"
	"game/status"
	"image"
	_ "image/png"
	"log"
	"math"
	"strings"
	"sync"
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
	SPEED               float64       = 150 //玩家走路移动速度
	SPEED_RUN           float64       = 200 //玩家跑步移动速度
	ISHIDDEN            uint8         = 1   //装备栏等隐藏标识
	ISITEMS             uint8         = 0   //物品和装备标识
	ISMINICOM           uint8         = 2   //MINI板标识
	ISNORCOM            uint8         = 3   //无标识 占位用
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

//Calculate Angle
func CaluteDirAtan2(x, y, x_tar, y_tar int64) float64 {
	dx := float64(x_tar - x)
	dy := float64(y_tar - y)
	j := math.Atan2(dy, dx) * 180 / math.Pi
	if j >= -270 && j < -90 {
		j = 270 + 180 + j
	} else {
		j = j + 90
	}
	return j
}

//Calculate Angle
func CaluteDisXY(dis, d float64) (int, int) {
	x := math.Sin(d*math.Pi/180) * dis
	y := math.Cos(d*math.Pi/180) * dis
	return int(math.Floor(395 + x)), int(math.Floor(240 - y))
}

//Calculate Direction
func CaluteDir(d float64) uint8 {
	if d >= 0 && d < 15 || d >= 345 && d < 360 {
		return 6
	}
	if d >= 15 && d < 35 {
		return 12
	}
	if d >= 35 && d < 55 {
		return 2
	}
	if d >= 55 && d < 75 {
		return 13
	}
	if d >= 75 && d < 105 {
		return 7
	}
	if d >= 105 && d < 125 {
		return 14
	}
	if d >= 125 && d < 145 {
		return 3
	}
	if d >= 145 && d < 165 {
		return 15
	}
	if d >= 165 && d < 195 {
		return 4
	}
	if d >= 195 && d < 215 {
		return 8
	}
	if d >= 215 && d < 235 {
		return 0
	}

	if d >= 235 && d < 255 {
		return 9
	}
	if d >= 255 && d < 285 {
		return 5
	}
	if d >= 285 && d < 305 {
		return 10
	}
	if d >= 305 && d < 325 {
		return 1
	}
	if d >= 325 && d < 345 {
		return 11
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

//计算屏幕坐标到世界坐标
func CalculateScreenToWorld(mx, my, px, py int) (int, int) {
	dx := mx - int(status.Config.PLAYERCENTERX)
	dy := my - int(status.Config.PLAYERCENTERY)
	return dx + px, dy + py
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
	yy := math.Floor((M_Plus_N - M_Minus_N) / 2)
	return int(xx), int(yy)
}

//根据瓦片坐标计算世界坐标
func TileToWorld(Tilex, Tiley int) (float64, float64) {
	return float64(3280 + (Tilex-Tiley)*80), float64((Tilex + Tiley) * 40)
}

//根据方向计算偏移距离
func CalculateSpeed(dir uint8, speed, dx, dy, s float64) (float64, float64) {
	// 1/60 is dt 1s 内刷新帧数
	moveX, moveY := 0.0, 0.0
	switch dir {
	case 0, 8, 9:
		moveX, moveY = -dx/s*speed/60, dy/s*speed/60
	case 1, 10, 11:
		moveX, moveY = -dx/s*speed/60, -dy/s*speed/60
	case 2, 12, 13:
		moveX, moveY = dx/s*speed/60, -dy/s*speed/60
	case 3, 14, 15:
		moveX, moveY = dx/s*speed/60, dy/s*speed/60
	case 4:
		moveX, moveY = 0, speed/60
	case 5:
		moveX, moveY = -speed/60, 0
	case 6:
		moveX, moveY = 0, -speed/60
	case 7:
		moveX, moveY = speed/60, 0
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
		box[0] = OffsetXY{0, 0}
		box[1] = OffsetXY{-4, -3}
		box[2] = OffsetXY{-45, -15}
		box[3] = OffsetXY{-15, -15}
	case "ba2":
		box[0] = OffsetXY{3, -7}
		box[1] = OffsetXY{8, -10}
		box[2] = OffsetXY{-55, -20}
		box[3] = OffsetXY{-10, -15}
	}
	return box
}

//消息对象复用
var PmPool = sync.Pool{
	New: func() interface{} { return new(pb.Message) },
}

var dataPool = sync.Pool{
	New: func() interface{} { return new(pb.Datas) },
}

//解包
func Unpack(msg []byte) *pb.Message {
	m := PmPool.Get().(*pb.Message)
	err := proto.Unmarshal(msg, m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}

//消息打包
func Pack(s bool, f, datas, msg string, p *pb.Player) []byte {
	m := PmPool.Get().(*pb.Message)
	pp := dataPool.Get().(*pb.Datas)
	m.Flag = f
	pp.Status = s
	pp.Data = datas
	pp.Mes = msg
	pp.Man = p
	m.Data = pp
	d, err := proto.Marshal(m)
	//放回对象池
	m.Reset()
	pp.Reset()
	dataPool.Put(pp)
	PmPool.Put(m)
	if err != nil {
		fmt.Println("has error###", err)
	}
	return d
}
