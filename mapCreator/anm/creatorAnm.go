package anm

import (
	"bytes"
	"embed"
	"fmt"
	"game/tools"
	"image"
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
	image    *embed.FS //静态资源获取
}

func NewAnm(images *embed.FS) *Anm {
	a := &Anm{
		anmiList: make([]*ebiten.Image, 0),
		wayList:  make([]*ebiten.Image, 0),
		huodui:   make([]*ebiten.Image, 0),
		op:       make([]*ebiten.DrawImageOptions, 0),
		image:    images,
	}
	return a
}
func (a *Anm) LoadAnm() {
	a.LoadXyList()
	for i := 0; i < 20; i++ {
		a.anmiList = append(a.anmiList, a.getImage(tools.ObjectPath+"/fire/frame_"+strconv.Itoa(i)+".png"))
		a.huodui = append(a.huodui, a.getImage(tools.ObjectPath+"/huodui/frame_"+strconv.Itoa(i)+".png"))
	}
	for i := 0; i < 15; i++ {
		a.wayList = append(a.wayList, a.getImage(tools.ObjectPath+"/waypoint/frame_"+strconv.Itoa(i)+".png"))
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
		xx, yy, _ := tools.GetCellXY(x, y)
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
	for i, k := range a.op {
		if i <= 14 {
			k.GeoM.Reset()
			k.GeoM.Translate(float64(a.xyPos[i].x)+offsetX, float64(a.xyPos[i].y)+offsetY)
			k.GeoM.Scale(Scale, Scale)
			screen.DrawImage(a.anmiList[frameIndex], k)
		}
	}

	a.op[15].GeoM.Reset()
	a.op[15].GeoM.Translate(float64(a.xyPos[15].x)+100+offsetX, float64(a.xyPos[15].y)+offsetY)
	a.op[15].GeoM.Scale(Scale, Scale)
	if frameIndex > 14 {
		screen.DrawImage(a.wayList[frameIndex-15], a.op[15])
	} else {
		screen.DrawImage(a.wayList[frameIndex], a.op[15])
	}

	a.op[16].GeoM.Reset()
	a.op[16].GeoM.Translate(float64(a.xyPos[16].x)+20+offsetX, float64(a.xyPos[16].y)-30+offsetY)
	a.op[16].GeoM.Scale(Scale, Scale)
	screen.DrawImage(a.huodui[frameIndex], a.op[16])

}

func (a *Anm) getImage(name string) *ebiten.Image {
	r, _ := a.image.ReadFile(name)
	m, _, err := image.Decode(bytes.NewReader(r))
	if err != nil {
		fmt.Println(err)
	}
	return ebiten.NewImageFromImage(m)
}
