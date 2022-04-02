package layout

import "github.com/hajimehoshi/ebiten/v2"

type icon struct {
	op     *ebiten.DrawImageOptions
	images *ebiten.Image
	flg    uint8
	f      func(i *icon)
}

func newIcon() *icon {
	i := &icon{
		op:  new(ebiten.DrawImageOptions),
		flg: 0,
	}
	i.op.Filter = ebiten.FilterLinear
	return i
}

//set images position
func (i *icon) pos(x, y float64) {
	i.op.GeoM.Translate(x, y)
}

//add imges
func (i *icon) addImage(m *ebiten.Image) {
	i.images = m
}

//register event to Ui
func (i *icon) addEvnet(fu func(i *icon)) {
	i.flg = 1
	i.f = fu
}
