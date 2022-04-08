package layout

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	op        *ebiten.DrawImageOptions
	images    *ebiten.Image
	hasEvent  uint8
	layer     uint8
	isDisplay bool
	f         func(i *Sprite)
	clickMinX int
	clickMinY int
	clickMaxX int
	clickMaxY int
	imagex    float64
	imagey    float64
}

//Create Icon Class
func newIcon() *Sprite {
	i := &Sprite{
		op:        new(ebiten.DrawImageOptions),
		hasEvent:  0,
		layer:     0,
		imagex:    0,
		imagey:    0,
		isDisplay: true,
	}
	i.op.Filter = ebiten.FilterLinear
	return i
}

//获取精灵的图片屏幕坐标
func (i *Sprite) GetPosition() (float64, float64) {
	return i.imagex, i.imagey
}

//Set Images Position
func (i *Sprite) SetPosition(x, y float64) {
	i.op.GeoM.Translate(x, y)
	i.imagex += x
	i.imagey += y
}

//Add Imges
func (i *Sprite) addImage(m *ebiten.Image) {
	i.images = m
}

//Register Event To Ui
func (i *Sprite) addEvent(fu func(i *Sprite)) {
	i.hasEvent = 1
	i.f = fu

}

func (i *Sprite) addEvnetRange(minX, minY, maxX, maxY int) {
	//Event range
	i.clickMinX = minX
	i.clickMinY = minY
	i.clickMaxX = maxX
	i.clickMaxY = maxY
}

//Quick Create icon
func QuickCreate(x, y float64, img *ebiten.Image, layer uint8, callBack func(i *Sprite), s ...int) *Sprite {
	op := newIcon()
	op.SetPosition(x, y)
	if len(s) == 4 {
		op.addEvnetRange(s[0], s[1], s[2], s[3])
	}
	if callBack != nil {
		op.addEvent(callBack)
	}
	op.layer = layer
	op.addImage(img)
	return op
}
