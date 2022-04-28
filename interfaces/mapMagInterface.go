package interfaces

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type MapInterface interface {
	LoadMap()
	ChangeMapTranslate(x, y float64)
	RenderFloor(screen *ebiten.Image, offsetX, offsetY float64)
	RenderWall(screen *ebiten.Image, offsetX, offsetY float64)
	GetBlock1Aera(x, y int) bool
	SortLayer(mapX, mapY int)
	RenderDropItems(screen *ebiten.Image, offsetX, offsetY float64, playX, playY float64)
	Render(screen *ebiten.Image, frameIndexFor20, frameIndexFor12 int, offsetX, offsetY float64)
	LoadAnm()
	PlayDropItemAnm(screen *ebiten.Image, x, y float64, name string, countsFor17 int) bool
	GetBlock1AeraUpdate(x, y int) bool
}
