package layout

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type icon struct {
	op        *ebiten.DrawImageOptions
	images    *ebiten.Image
	hasEvent  uint8
	layer     uint8
	isDisplay bool
	f         func(i *icon)
	clickMinX int
	clickMinY int
	clickMaxX int
	clickMaxY int
}

//Create Icon Class
func newIcon() *icon {
	i := &icon{
		op:        new(ebiten.DrawImageOptions),
		hasEvent:  0,
		layer:     0,
		isDisplay: true,
	}
	i.op.Filter = ebiten.FilterLinear
	return i
}

//Set Images Position
func (i *icon) pos(x, y float64) {
	i.op.GeoM.Translate(x, y)
}

//Add Imges
func (i *icon) addImage(m *ebiten.Image) {
	i.images = m
}

//Register Event To Ui
func (i *icon) addEvent(fu func(i *icon)) {
	i.hasEvent = 1
	i.f = fu

}

func (i *icon) addEvnetRange(minX, minY, maxX, maxY int) {
	//Event range
	i.clickMinX = minX
	i.clickMinY = minY
	i.clickMaxX = maxX
	i.clickMaxY = maxY
}

//Quick Create icon
func QuickCreate(x, y float64, img *ebiten.Image, layer uint8, callBack func(i *icon), s ...int) *icon {
	op := newIcon()
	op.pos(x, y)
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
