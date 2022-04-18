package anm

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
	x int
	y int
}
type Anm struct {
	anmiList []*ebiten.Image
	wayList  []*ebiten.Image
	huodui   []*ebiten.Image
	op       []*ebiten.DrawImageOptions
	xyPos    [17]postion
	image    *embed.FS            //静态资源获取
	status   *status.StatusManage //状态
}

func NewAnm(images *embed.FS, sta *status.StatusManage) *Anm {
	a := &Anm{
		anmiList: make([]*ebiten.Image, 0),
		wayList:  make([]*ebiten.Image, 0),
		huodui:   make([]*ebiten.Image, 0),
		op:       make([]*ebiten.DrawImageOptions, 0),
		image:    images,
		status:   sta,
	}
	return a
}
func (a *Anm) LoadAnm() {
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
}

//加载动画坐标
func (a *Anm) LoadXyList() {
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
		opr.GeoM.Translate(float64(xx), float64(yy))
		opr.GeoM.Scale(Scale, Scale)
		a.op = append(a.op, opr)
	}
}

//渲染object
func (a *Anm) Render(screen *ebiten.Image, frameIndex int, offsetX, offsetY float64) {
	//普通火台
	for i, k := range a.op {
		if i <= 14 {
			//shadow
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Rotate(-0.5)
			op.GeoM.Scale(1, 0.5)
			op.ColorM.Scale(0, 0, 0, 1)
			op.GeoM.Translate(float64(a.xyPos[i].x)+offsetX-50, float64(a.xyPos[i].y)+offsetY+50)
			screen.DrawImage(a.anmiList[frameIndex], op)
			//
			k.GeoM.Reset()
			k.GeoM.Translate(float64(a.xyPos[i].x)+offsetX, float64(a.xyPos[i].y)+offsetY)
			k.Filter = ebiten.FilterLinear
			k.GeoM.Scale(Scale, Scale)
			screen.DrawImage(a.anmiList[frameIndex], k)
		}
	}
	//传送点火焰
	a.op[15].GeoM.Reset()
	a.op[15].GeoM.Translate(float64(a.xyPos[15].x)+100+offsetX, float64(a.xyPos[15].y)+offsetY)
	a.op[15].Filter = ebiten.FilterLinear
	a.op[15].CompositeMode = ebiten.CompositeModeLighter
	a.op[15].GeoM.Scale(Scale, Scale)
	if frameIndex > 14 {
		screen.DrawImage(a.wayList[frameIndex-15], a.op[15])
	} else {
		screen.DrawImage(a.wayList[frameIndex], a.op[15])
	}
	//火堆
	a.op[16].GeoM.Reset()
	a.op[16].GeoM.Translate(float64(a.xyPos[16].x)+20+offsetX, float64(a.xyPos[16].y)-30+offsetY)
	a.op[16].Filter = ebiten.FilterLinear
	a.op[16].GeoM.Scale(Scale, Scale)
	screen.DrawImage(a.huodui[frameIndex], a.op[16])

}

//暗黑新手村专用 根据逻辑坐标 求具体坐标
func (a *Anm) GetCellXY(x, y int) (int, int, error) {
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
				return i*(-80) + sumX, startY + j*40, nil
			}

		}
	}
	str := "没有找到匹配的位置"
	return 0, 0, errors.New(str)
}
