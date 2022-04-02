package layout

import "github.com/hajimehoshi/ebiten/v2"

type icon struct {
	op     *ebiten.DrawImageOptions
	images *ebiten.Image
	flg    uint8
	f      func(i *icon)
}

//Create Icon Class
func newIcon() *icon {
	i := &icon{
		op:  new(ebiten.DrawImageOptions),
		flg: 0,
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
func (i *icon) addEvnet(fu func(i *icon)) {
	i.flg = 1
	i.f = fu
}
