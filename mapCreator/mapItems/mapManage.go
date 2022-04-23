package mapItems

import (
	"embed"
	"errors"
	"game/status"
	"game/tools"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

var Scale float64 = 1

type postion struct {
	x float64
	y float64
}

//掉落物品
type dropItem struct {
	name string
	pos  postion
}
type MapItems struct {
	anmiList      []*ebiten.Image //火把动画图集
	wayList       []*ebiten.Image //瞬间移动动画图集
	huodui        []*ebiten.Image //火堆动画图集
	NPC           []*ebiten.Image //NPC图集
	dropAnm       []*ebiten.Image //掉落动画图集
	dropItemsList []dropItem      //掉落物品一栏
	op            []*ebiten.DrawImageOptions
	xyPos         [17]postion
	image         *embed.FS            //静态资源获取
	status        *status.StatusManage //状态
}

func New(images *embed.FS, sta *status.StatusManage) *MapItems {
	a := &MapItems{
		anmiList:      make([]*ebiten.Image, 0),
		wayList:       make([]*ebiten.Image, 0),
		huodui:        make([]*ebiten.Image, 0),
		NPC:           make([]*ebiten.Image, 0),
		dropAnm:       make([]*ebiten.Image, 0),
		dropItemsList: make([]dropItem, 0),
		op:            make([]*ebiten.DrawImageOptions, 0),
		image:         images,
		status:        sta,
	}
	return a
}
func (a *MapItems) LoadAnm() {
	a.LoadXyList()
	for i := 0; i < 20; i++ {
		o, _ := a.image.ReadFile(tools.ObjectPath + "/fire/frame_" + strconv.Itoa(i) + ".png")
		a.anmiList = append(a.anmiList, tools.GetEbitenImage(o))
		o, _ = a.image.ReadFile(tools.ObjectPath + "/huodui/frame_" + strconv.Itoa(i) + ".png")
		a.huodui = append(a.huodui, tools.GetEbitenImage(o))
	}
	for i := 0; i < 15; i++ {
		o, _ := a.image.ReadFile(tools.ObjectPath + "/waypoint/frame_" + strconv.Itoa(i) + ".png")
		a.wayList = append(a.wayList, tools.GetEbitenImage(o))

	}
	for i := 0; i < 17; i++ {
		o, _ := a.image.ReadFile("resource/itemsdrop/c_" + strconv.Itoa(i) + ".png")
		a.dropAnm = append(a.dropAnm, tools.GetEbitenImage(o))
	}
	for i := 0; i < 12; i++ {
		o, _ := a.image.ReadFile("resource/monster/di/di_" + strconv.Itoa(i) + ".png")
		a.NPC = append(a.NPC, tools.GetEbitenImage(o))
	}
}

//加载动画坐标
func (a *MapItems) LoadXyList() {
	syList := [17]string{
		"22,24", "23,21", "22,16", "21,12", "24,10", "27,15", "33,12", "34,8", "38,10", "41,10", "43,8", "41,13", "46,15", "48,19", "44,24", "35,10", "31,16",
	}
	for i, k := range syList {
		re := strings.Split(k, ",")
		x, _ := strconv.Atoi(re[0])
		y, _ := strconv.Atoi(re[1])
		xx, yy, _ := a.GetCellXY(x, y)
		a.xyPos[i].x = xx
		a.xyPos[i].y = yy
		opr := &ebiten.DrawImageOptions{}
		opr.GeoM.Translate(xx, yy)
		opr.GeoM.Scale(Scale, Scale)
		a.op = append(a.op, opr)
	}
}

//渲染地图上物体
func (a *MapItems) Render(screen *ebiten.Image, frameIndexFor20, frameIndexFor12 int, offsetX, offsetY float64) {
	//普通火台
	for i, k := range a.op {
		if i <= 14 {
			//shadow
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Rotate(-0.5)
			op.GeoM.Scale(1, 0.5)
			op.ColorM.Scale(0, 0, 0, 1)
			op.GeoM.Translate(float64(a.xyPos[i].x)+offsetX-50, float64(a.xyPos[i].y)+offsetY+50)
			screen.DrawImage(a.anmiList[frameIndexFor20], op)
			//
			k.GeoM.Reset()
			k.GeoM.Translate(float64(a.xyPos[i].x)+offsetX, float64(a.xyPos[i].y)+offsetY)
			k.Filter = ebiten.FilterLinear
			k.GeoM.Scale(Scale, Scale)
			screen.DrawImage(a.anmiList[frameIndexFor20], k)
		}
	}
	//传送点火焰
	a.op[15].GeoM.Reset()
	a.op[15].GeoM.Translate(float64(a.xyPos[15].x)+100+offsetX, float64(a.xyPos[15].y)+offsetY)
	a.op[15].Filter = ebiten.FilterLinear
	a.op[15].CompositeMode = ebiten.CompositeModeLighter
	a.op[15].GeoM.Scale(Scale, Scale)
	if frameIndexFor20 > 14 {
		screen.DrawImage(a.wayList[frameIndexFor20-15], a.op[15])
	} else {
		screen.DrawImage(a.wayList[frameIndexFor20], a.op[15])
	}
	//火堆
	a.op[16].GeoM.Reset()
	a.op[16].GeoM.Translate(float64(a.xyPos[16].x)+20+offsetX, float64(a.xyPos[16].y)-30+offsetY)
	a.op[16].Filter = ebiten.FilterLinear
	a.op[16].GeoM.Scale(Scale, Scale)
	screen.DrawImage(a.huodui[frameIndexFor20], a.op[16])

	//掉落物品
	for i := 0; i < len(a.dropItemsList); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(a.dropItemsList[i].pos.x+offsetX, a.dropItemsList[i].pos.y+offsetY)
		screen.DrawImage(a.dropAnm[16], op)
	}

	//NPC
	//shadow
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(-0.5)
	op.GeoM.Scale(1, 0.5)
	op.ColorM.Scale(0, 0, 0, 1)
	op.GeoM.Translate(float64(a.xyPos[16].x)+50+offsetX, float64(a.xyPos[16].y)+60+offsetY)
	screen.DrawImage(a.NPC[frameIndexFor12], op)
	op1 := &ebiten.DrawImageOptions{}
	op1.GeoM.Translate(float64(a.xyPos[16].x)+100+offsetX, float64(a.xyPos[16].y)-35+offsetY)
	screen.DrawImage(a.NPC[frameIndexFor12], op1)
}

//暗黑新手村专用 根据逻辑坐标 求具体坐标
func (a *MapItems) GetCellXY(x, y int) (float64, float64, error) {
	if x < 0 || x > 57 || y < 0 || y > 41 {
		str := "坐标范围不正确"
		return 0, 0, errors.New(str)
	}
	startY := 0
	sumX := 0
	for i := 0; i < a.status.ReadMapSizeHeight; i++ {
		if i > 0 {
			startY += 40
		}
		sumX = 0
		for j := 0; j < a.status.ReadMapSizeWidth; j++ {
			if j > 0 {
				sumX += 80
			}
			if j == x && y == i {
				return float64(i*(-80) + sumX), float64(startY + j*40), nil
			}

		}
	}
	str := "没有找到匹配的位置"
	return 0, 0, errors.New(str)
}

//播放丢物品动画
func (a *MapItems) PlayDropItemAnm(screen *ebiten.Image, x, y int) {
	go func() {
		countForMap := 0
		countsForMap := 0
		op := &ebiten.DrawImageOptions{}
		op.Filter = ebiten.FilterLinear
		op.GeoM.Translate(float64(tools.LAYOUTX/2-220), float64(tools.LAYOUTY/2)-50)
		for {
			countForMap++
			screen.DrawImage(a.dropAnm[countsForMap], op)
			//切换图
			if countForMap > 5990 {
				countsForMap++
				countForMap = 0
				if countsForMap >= 16 {
					a.InsertOnLoadItesm("dun", x, y)
					return
				}
			}
		}
	}()
}

//记录掉落在地面的物品信息
func (a *MapItems) InsertOnLoadItesm(name string, x, y int) {
	var i dropItem
	i.name = name
	xx, yy, _ := a.GetCellXY(x-1, y-1)
	i.pos.x = xx
	i.pos.y = yy
	a.dropItemsList = append(a.dropItemsList, i)
}
