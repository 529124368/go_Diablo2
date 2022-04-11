package layout

import "github.com/hajimehoshi/ebiten/v2"

//精灵接口
type spriteInterface interface {
	GetPosition() (float64, float64)
	GetSpriteSize() (int, int)
	SetPosition(x, y float64)
	addImage(m *ebiten.Image)
	addEvent(fu func(s spriteInterface))
	addClickRange()
	DrawItemsBgByCustom(positionX, postionY float64, width, height int, screen *ebiten.Image)
	QuickDrawItemsBg(screen *ebiten.Image)
}
