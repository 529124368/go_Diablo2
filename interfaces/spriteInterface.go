package interfaces

import "github.com/hajimehoshi/ebiten/v2"

//精灵接口
type SpriteInterface interface {
	GetPosition() (float64, float64)
	GetSpriteSize() (int, int)
	SetPosition(x, y float64)
	AddImage(m string)
	AddEvent(fu func(s SpriteInterface))
	AddClickRange()
	DrawItemsBgByCustom(positionX, postionY float64, width, height int, screen *ebiten.Image)
	QuickDrawItemsBg(screen *ebiten.Image)
}
